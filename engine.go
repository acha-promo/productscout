package core

import (
	"errors"
	"fmt"
	"sync"
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
		Scrape(gtin string) ([]Product, error)
		Info() Website
	}

	Config struct {
		Debug  bool
		Logger Logger
	}

	Engine struct {
		Config   Config
		Scrapers []Scraper
	}
)

var (
	ErrMissingScraper = errors.New("missing scraper")
)

func (e *Engine) Search(gtin string) ([]Product, error) {

	if e.Config.Logger == nil {
		e.Config.Logger = &DefaultLogger{}
	}

	var wg sync.WaitGroup
	results := []Product{}
	resultsMutex := &sync.Mutex{}

	if len(e.Scrapers) == 0 {
		return results, ErrMissingScraper
	}

	for _, scraper := range e.Scrapers {
		e.debug("Scraping", scraper.Info().URL)
		wg.Add(1)
		go func(s Scraper) {
			defer wg.Done()
			products, err := s.Scrape(gtin)
			if err == nil {
				resultsMutex.Lock()
				results = append(results, products...)
				resultsMutex.Unlock()
			} else {
				e.Config.Logger.Error(fmt.Errorf("scraper %T failed: %w", s, err))
			}
		}(scraper)
	}

	wg.Wait()

	return results, nil
}

func (e *Engine) debug(args ...any) {
	if e.Config.Debug {
		e.Config.Logger.Debug(args...)
	}
}
