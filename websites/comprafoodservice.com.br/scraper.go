package comprafoodservicecombr

/*
curl 'https://api.linximpulse.com/engage/search/v3/search?apiKey=cfs-new&page=1&resultsPerPage=40&terms=7898422745523&sortBy=relevance&salesChannel=default' \
  -H 'authority: api.linximpulse.com' \
  -H 'origin: https://www.comprafoodservice.com.br' \
  -H 'referer: https://www.comprafoodservice.com.br/buscar/' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36'
*/

import (
	"fmt"

	core "achapromo.com/productscout"
	"achapromo.com/productscout/httpclient"
)

const (
	URL = "https://api.linximpulse.com"
)

type (
	Scraper struct {
		HttpClient *httpclient.HttpClient
	}

	ProductData struct {
		RequestId string `json:"requestId"`
		SearchId  string `json:"searchId"`
		Filters   []struct {
			Id        int    `json:"id"`
			Attribute string `json:"attribute"`
			Type      string `json:"type"`
			FType     int    `json:"fType"`
			Values    []struct {
				Label   string `json:"label"`
				Size    int    `json:"size"`
				IdO     string `json:"idO"`
				Id      int    `json:"id"`
				Filters []struct {
					Label     string `json:"label"`
					Size      int    `json:"size"`
					IdO       string `json:"idO"`
					Id        int    `json:"id"`
					ApplyLink string `json:"applyLink"`
				} `json:"filters"`
				ApplyLink string `json:"applyLink"`
			} `json:"values"`
		} `json:"filters"`
		Size       int `json:"size"`
		Pagination struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"pagination"`
		Products []struct {
			Id          string `json:"id"`
			CollectInfo struct {
				ProductId string        `json:"productId"`
				SkuList   []interface{} `json:"skuList"`
			} `json:"collectInfo"`
			ClickUrl string `json:"clickUrl"`
			Name     string `json:"name"`
			Price    int    `json:"price"`
			OldPrice int    `json:"oldPrice"`

			Url    string `json:"url"`
			Images struct {
				Default string `json:"default"`
			} `json:"images"`
			Installment    []interface{} `json:"installment"`
			CustomBusiness struct {
				Search struct {
					Snippets []struct {
						Term   string   `json:"term"`
						Fields []string `json:"fields"`
					} `json:"snippets"`
				} `json:"search"`
			} `json:"customBusiness"`
			Status     string `json:"status"`
			Categories []struct {
				Id      string   `json:"id"`
				Name    string   `json:"name"`
				Parents []string `json:"parents"`
			} `json:"categories"`
			Tags []struct {
				Id      string   `json:"id"`
				Name    string   `json:"name"`
				Parents []string `json:"parents"`
			} `json:"tags"`
			Specs struct {
				Caixa []struct {
					Id         string `json:"id"`
					Label      string `json:"label"`
					Properties struct {
					} `json:"properties"`
				} `json:"caixa"`
			} `json:"specs"`
			Created     string `json:"created"`
			Brand       string `json:"brand"`
			SelectedSku string `json:"selectedSku"`
			CId         string `json:"cId"`
			IId         int    `json:"iId"`
			Skus        []struct {
				Sku   string `json:"sku"`
				Specs struct {
					Caixa string `json:"caixa"`
				} `json:"specs"`
				Properties struct {
					Status   string `json:"status"`
					Price    int    `json:"price"`
					OldPrice int    `json:"oldPrice"`
					Images   struct {
						Default string `json:"default"`
					} `json:"images"`
					Created string `json:"created"`
					Details struct {
						Caixa       string   `json:"caixa"`
						Marca       string   `json:"marca"`
						SubSkus     []string `json:"subSkus"`
						Label       string   `json:"label"`
						Peso        string   `json:"peso"`
						Embalagem   string   `json:"embalagem"`
						Brand       string   `json:"brand"`
						Prodcat     string   `json:"prodcat"`
						Mpn         string   `json:"mpn"`
						Bu          string   `json:"bu"`
						Measurement struct {
						} `json:"measurement"`
						Url string `json:"url"`
					} `json:"details"`
				} `json:"properties"`
			} `json:"skus"`
			Details struct {
				Embalagem    []string `json:"embalagem"`
				SalesChannel []string `json:"salesChannel"`
				CategoryName []string `json:"categoryName"`
			} `json:"details"`
		} `json:"products"`
		Sort []struct {
			Label     string `json:"label"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			ApplyLink string `json:"applyLink"`
		} `json:"sort"`
		Queries struct {
			Original   string `json:"original"`
			Normalized string `json:"normalized"`
			Processed  string `json:"processed"`
			QueryType  string `json:"queryType"`
		} `json:"queries"`
	}
)

func (s *Scraper) Scrape(gtin string) ([]core.Product, error) {
	productData, err := s.fetchProductData(gtin)
	if err != nil {
		return nil, err
	}

	if len(productData.Products) == 0 {
		return nil, core.ErrProductNotFound
	}

	products := []core.Product{}
	for _, productData := range productData.Products {
		product := core.Product{
			Name: productData.Name,
			GTIN: productData.Id,
			URL:  url(gtin),
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *Scraper) Info() core.Website {
	return core.Website{URL: URL}
}

// func sanitize(name string) string {
// 	if i := strings.Index(name, " (from barcode.monster)"); i != -1 {
// 		name = name[:i]
// 	}
// 	return name
// }

func (s *Scraper) fetchProductData(gtin string) (*ProductData, error) {

	s.HttpClient.SetHeaders(map[string]string{
		"authority":  "api.linximpulse.com",
		"origin":     "https://www.comprafoodservice.com.br",
		"referer":    "https://www.comprafoodservice.com.br/buscar/",
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	})

	var productData ProductData
	err := s.HttpClient.GetJSON(url(gtin), &productData)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}

func url(product string) string {
	return fmt.Sprintf("%s/engage/search/v3/search?apiKey=cfs-new&page=1&resultsPerPage=40&terms=%s&sortBy=relevance&salesChannel=default", URL, product)
}
