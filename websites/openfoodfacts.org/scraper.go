package openfoodfactsorg

import (
	"fmt"

	core "achapromo.com/gtinscout"
)

const (
	URL = "https://world.openfoodfacts.org"
)

type (
	Scraper struct {
		HttpClient *core.HttpClient
	}

	Data struct {
		Status  int    `json:"status"`
		Code    string `json:"code"`
		Product `json:"product"`
	}

	Product struct {
		NamePt       string `json:"product_name_pt"`
		Quantity     string `json:"product_quantity"`
		QuantityUnit string `json:"product_quantity_unit"`
		Packaging    string `json:"packaging"`
	}
)

func (s *Scraper) Scrape(gtin string) ([]core.Product, error) {
	data, err := s.fetchProductData(gtin)
	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, core.ErrProductNotFound
	}

	product := core.Product{
		Name: concat(
			data.Product.NamePt,
			data.Product.Packaging,
			data.Product.Quantity,
			data.Product.QuantityUnit,
		),
		GTIN: data.Code,
		URL:  url(gtin),
	}

	return []core.Product{product}, nil
}

func (s *Scraper) Info() core.Website {
	return core.Website{URL: URL}
}

func concat(name, packaging, quantity, unit string) string {
	return fmt.Sprintf("%s %s %s%s", name, packaging, quantity, unit)
}

func (s *Scraper) fetchProductData(gtin string) (*Data, error) {
	var productData Data
	err := s.HttpClient.GetJSON(url(gtin), &productData)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}

func url(gtin string) string {
	return fmt.Sprintf("%s/api/v0/product/%s.json", URL, gtin)
}
