package barcodemonster_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"achapromo.com/gtinscout/httpclient"
	barcodemonster "achapromo.com/gtinscout/websites/barcode.monster"
	"github.com/stretchr/testify/assert"
)

func TestScraper_ScrapeProductData_Valid(t *testing.T) {
	expectedURL := "/api/7898422745523"
	expectedResponse := `{
		class: "EAN13",
		code: "7898422745523",
		description: "Sabao em po ala lavanda 500g (from barcode.monster)",
		image_url: "https://centralmidia-riograndense.s3-sa-east-1.amazonaws.com/550/7898422745523_1.jpg",
		size: "",
		status: "active"
		}`

	server := createServer(t, expectedURL, expectedResponse)
	defer server.Close()
	scraper := barcodemonster.Scraper{HttpClient: httpclient.NewHttpClient()}

	// Call
	productData, err := scraper.Scrape("7898422745523")

	// Verify
	assert.NoError(t, err)
	assert.Len(t, productData, 1)
	if len(productData) == 0 {
		return
	}
	product := productData[0]
	assert.Equal(t, "Sabao em po ala lavanda 500g", product.Name)
	assert.Equal(t, "7898422745523", product.GTIN)
}

func createServer(t *testing.T, expectedURL, expectedResponse string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, expectedURL, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(expectedResponse))
	}))
	return server
}
