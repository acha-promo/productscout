package paguemenoscombr

import (
	"fmt"

	core "achapromo.com/productscout"
	"achapromo.com/productscout/httpclient"
)

const (
	URL = "https://prod.apipmenos.com"
)

type (
	Scraper struct {
		HttpClient *httpclient.HttpClient
	}

	Result struct {
		RequestId string `json:"requestId"`
		SearchId  string `json:"searchId"`
		Filters   []struct {
			ID        int    `json:"id"`
			Attribute string `json:"attribute"`
			Type      string `json:"type"`
			FType     int    `json:"fType"`
			Selected  bool   `json:"selected"`
			Values    []struct {
				Label      string      `json:"label"`
				Size       int         `json:"size"`
				IDO        interface{} `json:"idO"`
				ID         int         `json:"id"`
				Filters    interface{} `json:"filters"`
				Selected   bool        `json:"selected"`
				RemoveLink interface{} `json:"removeLink"`
				ApplyLink  interface{} `json:"applyLink"`
				UnityId    interface{} `json:"unityId"`
				UN         interface{} `json:"unN"`
				Min        interface{} `json:"min"`
				Max        interface{} `json:"max"`
			} `json:"values"`
			Label      interface{} `json:"label"`
			Size       int         `json:"size"`
			IDO        interface{} `json:"idO"`
			RemoveLink interface{} `json:"removeLink"`
			ApplyLink  interface{} `json:"applyLink"`
			Filters    interface{} `json:"filters"`
		} `json:"filters"`
		RemoveAllFiltersLink interface{}   `json:"removeAllFiltersLink"`
		SelectedFilters      []interface{} `json:"selectedFilters"`
		Size                 int           `json:"size"`
		Pagination           struct {
			First    string      `json:"first"`
			Last     string      `json:"last"`
			Prev     interface{} `json:"prev"`
			Next     interface{} `json:"next"`
			LastPage bool        `json:"lastPage"`
		} `json:"pagination"`
		Suggestions  interface{}   `json:"suggestions"`
		Banners      interface{}   `json:"banners"`
		QuickFilters []interface{} `json:"quickFilters"`
		Sort         []struct {
			Label      string      `json:"label"`
			Name       string      `json:"name"`
			Type       string      `json:"type"`
			ApplyLink  string      `json:"applyLink"`
			RemoveLink interface{} `json:"removeLink"`
			Selected   bool        `json:"selected"`
		} `json:"sort"`
		Queries struct {
			Original   string `json:"original"`
			Normalized string `json:"normalized"`
			Processed  string `json:"processed"`
			QueryType  string `json:"queryType"`
		} `json:"queries"`
		UltimaPagina  bool        `json:"ultimaPagina"`
		TotalProdutos string      `json:"totalProdutos"`
		IsVitrine     bool        `json:"isVitrine"`
		Collection    interface{} `json:"collection"`
		Headline      struct {
			Titulo    interface{} `json:"titulo"`
			SubTitulo interface{} `json:"subTitulo"`
		} `json:"headline"`
		IsCategory interface{} `json:"isCategory"`
		QueryType  interface{} `json:"queryType"`
		Link       interface{} `json:"link"`
		Data       []struct {
			ProductID        int         `json:"productId"`
			ClickURL         string      `json:"clickUrl"`
			URL              string      `json:"url"`
			TrackingID       string      `json:"trackingId"`
			ProductName      string      `json:"productName"`
			BrandID          int         `json:"brandId"`
			Brand            string      `json:"brand"`
			CategoryID       interface{} `json:"categoryId"`
			Description      string      `json:"description"`
			IsProdutoDigital bool        `json:"isProdutoDigital"`
			Items            []struct {
				SKU         int    `json:"sku"`
				ItemID      int    `json:"itemId"`
				EAN         string `json:"ean"`
				ReferenceID []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"referenceId"`
				Images []struct {
					ImageID  int    `json:"imageId"`
					ImageTag string `json:"imageTag"`
					ImageURL string `json:"imageUrl"`
				} `json:"images"`
				Sellers []struct {
					SellerID        string `json:"sellerId"`
					SellerName      string `json:"sellerName"`
					CommertialOffer struct {
						DeliverySLASamplesPerRegion interface{}   `json:"deliverySlaSamplesPerRegion"`
						Price                       float64       `json:"price"`
						ListPrice                   float64       `json:"listPrice"`
						PriceWithoutDiscount        float64       `json:"priceWithoutDiscount"`
						Discount                    int           `json:"discount"`
						DiscountLabel               string        `json:"discountLabel"`
						DiscountHighLightApp        bool          `json:"discountHighLightApp"`
						DiscountHighLight           []interface{} `json:"discountHighLight"`
						GiftSKUIds                  []interface{} `json:"giftSkuIds"`
						Teasers                     []struct {
							NameKBackingField       string `json:"nameKBackingField"`
							Value                   string `json:"value"`
							ConditionsKBackingField struct {
								MinimumQuantityKBackingField string `json:"minimumQuantityKBackingField"`
								Value                        string `json:"value"`
							} `json:"conditionsKBackingField"`
							Conditions struct {
								MinimumQuantityKBackingField string `json:"minimumQuantityKBackingField"`
								Value                        string `json:"value"`
							} `json:"conditions"`
						} `json:"teasers"`
						BuyTogether            []interface{} `json:"buyTogether"`
						ItemMetadataAttachment []interface{} `json:"itemMetadataAttachment"`
						GetInfoErrorMessage    interface{}   `json:"getInfoErrorMessage"`
					} `json:"commertialOffer"`
				} `json:"sellers"`
			} `json:"items"`
			Medicamento           []string      `json:"medicamento"`
			MedicamentoControlado []string      `json:"medicamentoControlado"`
			MarcaPropria          []string      `json:"marcaPropria"`
			Generico              []string      `json:"generico"`
			AME                   []interface{} `json:"ame"`
			Antibiotico           []string      `json:"antibiotico"`
			PBM                   []interface{} `json:"pbm"`
			AdministradoraPBM     []interface{} `json:"administradoraPBM"`
			ProgramaPBM           []interface{} `json:"programaPBM"`
			MenorPrecoPBM         []interface{} `json:"menorPrecoPBM"`
			EanPBM                []interface{} `json:"eanPBM"`
			LinkLogoProgramaPBM   []interface{} `json:"linkLogoProgramaPBM"`
			PontosProgramaPBM     []interface{} `json:"pontosProgramaPBM"`
			ProgramaDermoPBM      []interface{} `json:"programaDermoPBM"`
			ApresentarPrecoDePor  []string      `json:"apresentarPrecoDePor"`
			DescontoPBM           []interface{} `json:"descontoPbm"`
			DescricaoBula         []interface{} `json:"descricaoBula"`
			Categories            []string      `json:"categories"`
			CategoriesIds         []string      `json:"categoriesIds"`
			MetaTagDescription    string        `json:"metaTagDescription"`
			MedAme                bool          `json:"medAme"`
			Cashback              []string      `json:"cashback"`
		} `json:"data"`
	}
)

func (s *Scraper) Scrape(gtin string) ([]core.Product, error) {
	res, err := s.fetchProductData(gtin)
	if err != nil {
		return nil, err
	}

	if len(res.Data) == 0 {
		return nil, core.ErrProductNotFound
	}

	products := []core.Product{}
	for _, product := range res.Data {

		gtin := ""
		if len(product.Items) > 0 {
			gtin = product.Items[0].EAN
		}
		products = append(products, core.Product{
			Name: product.ProductName,
			GTIN: gtin,
			URL:  url(gtin),
		})
	}

	return products, nil
}

func (s *Scraper) Info() core.Website {
	return core.Website{URL: URL}
}

func (s *Scraper) fetchProductData(gtin string) (*Result, error) {
	var productData Result

	s.HttpClient.SetHeaders(map[string]string{
		"sec-ch-ua":          `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`,
		"Accept":             "application/json, text/plain, */*",
		"Referer":            "https://www.paguemenos.com.br/",
		"sec-ch-ua-mobile":   "?0",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		"x-api-key":          "Yt1tDH9WNx5pmTXrBPBFH8mHAMJ5Gbb3dbSdu12d",
		"sec-ch-ua-platform": "macOS",
	})

	err := s.HttpClient.GetJSON(url(gtin), &productData)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}

func url(product string) string {
	return fmt.Sprintf("%s/buscacatalogo/api/searchurl?salesChannel=1&company=1&url=%%2Fengage%%2Fsearch%%2Fv3%%2Fsearch%%3Fapikey%%3Dfarmacia-paguemenos%%26terms%%3D%s%%26page%%3D1%%26resultsperpage%%3D48%%26showonlyavailable%%3Dfalse%%26allowredirect%%3Dtrue&deviceId=47298a9e-9593-4895-87bf-e05ac4c3b24f&source=desktop", URL, product)
}
