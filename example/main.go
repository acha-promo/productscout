package main

import (
	"fmt"
	"time"

	core "achapromo.com/gtinscout"
	barcodemonster "achapromo.com/gtinscout/websites/barcode.monster"
)

func main() {

	httpClient := core.NewHttpClient(
		core.WithUserAgent("gtinscout"),
		core.WithTimeout(time.Second*30),
	)

	engine := core.Engine{
		Config: core.Config{
			Debug: true,
		},
		Scrapers: []core.Scraper{
			&barcodemonster.Scraper{HttpClient: httpClient},
		},
	}

	products, err := engine.Search("7898422745523")
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("%+v\n", product)
	}

}
