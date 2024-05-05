package barcodemonster

import (
	"fmt"
	"strings"

	core "achapromo.com/productscout"
	"achapromo.com/productscout/httpclient"
)

const (
	URL = "https://barcode.monster"
)

type (
	Scraper struct {
		HttpClient *httpclient.HttpClient
	}

	ProductData struct {
		Class       string `json:"class"`
		Code        string `json:"code"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		Size        string `json:"size"`
		Status      string `json:"status"`
	}
)

func (s *Scraper) Scrape(gtin string) ([]core.Product, error) {
	productData, err := s.fetchProductData(gtin)
	if err != nil {
		return nil, err
	}

	if productData.Status == "not found" {
		return nil, core.ErrProductNotFound
	}

	product := core.Product{
		Name: sanitize(productData.Description),
		GTIN: productData.Code,
		URL:  url(gtin),
	}

	return []core.Product{product}, nil
}

func (s *Scraper) Info() core.Website {
	return core.Website{URL: URL}
}

func sanitize(name string) string {
	if i := strings.Index(name, " (from barcode.monster)"); i != -1 {
		name = name[:i]
	}
	return name
}

func (s *Scraper) fetchProductData(gtin string) (*ProductData, error) {
	var productData ProductData
	err := s.HttpClient.GetJSON(url(gtin), &productData)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}

func url(gtin string) string {
	return fmt.Sprintf("%s/api/%s", URL, gtin)
}
