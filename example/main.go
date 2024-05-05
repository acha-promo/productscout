package main

import (
	"fmt"
	"time"

	core "achapromo.com/productscout"
	"achapromo.com/productscout/httpclient"
	barcodemonster "achapromo.com/productscout/websites/barcode.monster"
	comprafoodservicecombr "achapromo.com/productscout/websites/comprafoodservice.com.br"
	openfoodfactsorg "achapromo.com/productscout/websites/openfoodfacts.org"
)

func main() {

	httpClient := httpclient.NewHttpClient(
		httpclient.WithUserAgent("gtinscout"),
		httpclient.WithTimeout(time.Second*30),
	)

	engine := core.NewEngine(
		core.WithDebug(),
		core.WithMaxConcurrency(20),
		core.WithScrapers(
			&comprafoodservicecombr.Scraper{HttpClient: httpClient},
			&barcodemonster.Scraper{HttpClient: httpClient},
			&openfoodfactsorg.Scraper{HttpClient: httpClient},
		),
	)

	products, err := engine.Search("7898215151784")
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("%+v\n", product)
	}
}
