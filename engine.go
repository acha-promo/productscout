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
		config   Config
		scrapers []Scraper
	}

	// Option configures the Engine
	Option func(*Engine)
)

var (
	ErrMissingScraper  = errors.New("missing scraper")
	ErrProductNotFound = errors.New("product not found")

	gtinRegex = regexp.MustCompile(`^\d{13}$`)
)

func NewEngine(opts ...Option) *Engine {
	e := &Engine{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithTimeout sets the timeout for the Engine as a whole operation
func WithTimeout(d time.Duration) Option {
	return func(e *Engine) {
		e.config.Timeout = d
	}
}

// WithLogger sets the logger for the Engine
func WithLogger(l Logger) Option {
	return func(e *Engine) {
		e.config.Logger = l
	}
}

// WithDebug sets the debug mode for the Engine
func WithDebug() Option {
	return func(e *Engine) {
		e.config.Debug = true
	}
}

func WithMaxConcurrency(n int) Option {
	return func(e *Engine) {
		e.config.MaxConcurrency = n
	}
}

func WithScrapers(s ...Scraper) Option {
	return func(e *Engine) {
		e.scrapers = append(e.scrapers, s...)
	}
}

func (e *Engine) setup() {
	if e.config.Logger == nil {
		e.config.Logger = &DefaultLogger{}
	}

	if e.config.MaxConcurrency == 0 {
		e.config.MaxConcurrency = 10
	}
}

func (e *Engine) Search(product string) ([]Product, error) {
	e.setup()

	if len(e.scrapers) == 0 {
		return nil, ErrMissingScraper
	}

	var ctx context.Context
	var cancel context.CancelFunc
	if e.config.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), e.config.Timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Ensure all paths cancel the context to avoid context leak

	var wg sync.WaitGroup
	results := []Product{}
	resultsMutex := &sync.Mutex{}

	// Create a buffered channel to limit the number of goroutines
	semaphore := make(chan struct{}, e.config.MaxConcurrency)

	for idx, scraper := range e.scrapers {
		e.debug("Scraping", scraper.Info().URL)

		select {
		case semaphore <- struct{}{}: // Block if there are already MaxConcurrency goroutines running
		case <-ctx.Done(): // Check if the context deadline has been reached
			e.config.Logger.Error(fmt.Sprintf("timeout reached before starting all scrapers (exec: %d | non-exec: %d)", idx+1, len(e.scrapers[idx:])))
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
					e.config.Logger.Error(fmt.Errorf("scraper %T failed: %w", s, err))
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

func (e *Engine) AddScraper(s Scraper) {
	e.scrapers = append(e.scrapers, s)
}

func (e *Engine) debug(args ...any) {
	if e.config.Debug {
		e.config.Logger.Debug(args...)
	}
}

func isGTIN(s string) bool {
	return gtinRegex.MatchString(s)
}
