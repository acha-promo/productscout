package core

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sync"
	"time"
)

type (
	Product struct {
		Name string
		GTIN string
		URL  string
	}

	Website struct {
		URL string
	}

	Scraper interface {
		Scrape(product string) ([]Product, error)
		Info() Website
	}

	Config struct {
		Debug          bool
		Logger         Logger
		MaxConcurrency int
		Timeout        time.Duration
	}

	Engine struct {
		Config   Config
		Scrapers []Scraper
	}
)

var (
	ErrMissingScraper  = errors.New("missing scraper")
	ErrProductNotFound = errors.New("product not found")
)

func (e *Engine) init() {
	if e.Config.Logger == nil {
		e.Config.Logger = &DefaultLogger{}
	}

	if e.Config.MaxConcurrency == 0 {
		e.Config.MaxConcurrency = 10
	}
}

func (e *Engine) Search(product string) ([]Product, error) {
	e.init()

	if len(e.Scrapers) == 0 {
		return nil, ErrMissingScraper
	}

	var ctx context.Context
	var cancel context.CancelFunc
	if e.Config.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), e.Config.Timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Ensure all paths cancel the context to avoid context leak

	var wg sync.WaitGroup
	results := []Product{}
	resultsMutex := &sync.Mutex{}

	// Create a buffered channel to limit the number of goroutines
	semaphore := make(chan struct{}, e.Config.MaxConcurrency)

	for idx, scraper := range e.Scrapers {
		e.debug("Scraping", scraper.Info().URL)

		select {
		case semaphore <- struct{}{}: // Block if there are already MaxConcurrency goroutines running
		case <-ctx.Done(): // Check if the context deadline has been reached
			e.Config.Logger.Error(fmt.Sprintf("timeout reached before starting all scrapers (exec: %d | non-exec: %d)", idx+1, len(e.Scrapers[idx:])))
			return results, ctx.Err()
		}

		wg.Add(1)
		go func(s Scraper) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release the spot in the semaphore when the goroutine completes

			if ctx.Err() == nil { // Check context error to avoid executing if already timed out
				products, err := s.Scrape(product)
				if err == nil {
					resultsMutex.Lock()
					results = append(results, products...)
					resultsMutex.Unlock()
				} else {
					e.Config.Logger.Error(fmt.Errorf("scraper %T failed: %w", s, err))
				}
			}
		}(scraper)
	}

	wg.Wait() // Wait for all goroutines to finish

	if isGTIN(product) {
		filteredResults := make([]Product, 0)
		for _, result := range results {
			if result.GTIN == product {
				filteredResults = append(filteredResults, result)
			}
		}
		results = filteredResults
	}

	return results, nil
}

func (e *Engine) debug(args ...any) {
	if e.Config.Debug {
		e.Config.Logger.Debug(args...)
	}
}

func isGTIN(s string) bool {
	return regexp.MustCompile(`^\d{13}$`).MatchString(s)
}
