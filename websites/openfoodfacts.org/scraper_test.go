package openfoodfactsorg_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	core "achapromo.com/gtinscout"
	openfoodfactsorg "achapromo.com/gtinscout/websites/openfoodfacts.org"

	"github.com/stretchr/testify/assert"
)

func TestScraper_ScrapeProductData_Valid(t *testing.T) {
	expectedURL := "/api/v0/product/7898215151784.json"
	expectedResponse := getValidJson()

	server := createServer(t, expectedURL, expectedResponse)
	defer server.Close()
	scraper := openfoodfactsorg.Scraper{HttpClient: core.NewHttpClient()}

	// Call
	productData, err := scraper.Scrape("7898215151784")

	// Verify
	assert.NoError(t, err)
	assert.Len(t, productData, 1)
	if len(productData) == 0 {
		return
	}
	product := productData[0]
	assert.Equal(t, "Creme de Leite Piracanjuba Caixinha 200g", product.Name)
	assert.Equal(t, "7898215151784", product.GTIN)
}

func createServer(t *testing.T, expectedURL, expectedResponse string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, expectedURL, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(expectedResponse))
	}))
	return server
}

func getValidJson() string {
	return `{
		code: "7898215151784",
		product: {
		_id: "7898215151784",
		_keywords: [
		"cream",
		"creme",
		"dairie",
		"de",
		"fsc",
		"leite",
		"milk",
		"piracanjuba",
		"uht",
		"uht-milk",
		"whole"
		],
		added_countries_tags: [ ],
		additives_n: 0,
		additives_old_n: 0,
		additives_old_tags: [ ],
		additives_original_tags: [ ],
		additives_tags: [ ],
		allergens: "",
		allergens_from_ingredients: "en:milk, ALÉRGICOS: CONTÉM LEITE E DERIVADOS. CONTÉM LACTOSE. NÃO CONTÉM GLÚTEN., CONTÉM LEITE E DERIVADOS, CONTÉM LACTOSE, CONTÉM LEITE E DERIVADOS, CONTÉM LACTOSE",
		allergens_from_user: "(pt) ",
		allergens_hierarchy: [
		"en:milk",
		"pt:ALÉRGICOS: CONTÉM LEITE E DERIVADOS. CONTÉM LACTOSE. NÃO CONTÉM GLÚTEN."
		],
		allergens_tags: [
		"en:milk",
		"pt:alergicos-contem-leite-e-derivados-contem-lactose-nao-contem-gluten"
		],
		amino_acids_tags: [ ],
		brands: "Piracanjuba",
		brands_tags: [
		"piracanjuba"
		],
		categories: "Dairies, Creams, Milks, UHT Creams, Whole milks, en:uht-milks",
		categories_hierarchy: [
		"en:dairies",
		"en:creams",
		"en:milks",
		"en:homogenized-milks",
		"en:uht-milks",
		"en:uht-creams",
		"en:whole-milks"
		],
		categories_lc: "en",
		categories_old: "Dairies, Creams, UHT Creams, en:milks",
		categories_properties: {
		agribalyse_proxy_food_code:en: "19024",
		ciqual_food_code:en: "19039"
		},
		categories_properties_tags: [
		"all-products",
		"categories-known",
		"agribalyse-food-code-unknown",
		"agribalyse-proxy-food-code-19024",
		"agribalyse-proxy-food-code-known",
		"ciqual-food-code-19039",
		"ciqual-food-code-known",
		"agribalyse-known",
		"agribalyse-19024"
		],
		categories_tags: [
		"en:dairies",
		"en:creams",
		"en:milks",
		"en:homogenized-milks",
		"en:uht-milks",
		"en:uht-creams",
		"en:whole-milks"
		],
		category_properties: { },
		checkers_tags: [ ],
		ciqual_food_name_tags: [
		"unknown"
		],
		cities_tags: [ ],
		code: "7898215151784",
		codes_tags: [
		"code-13",
		"7898215151xxx",
		"789821515xxxx",
		"78982151xxxxx",
		"7898215xxxxxx",
		"789821xxxxxxx",
		"78982xxxxxxxx",
		"7898xxxxxxxxx",
		"789xxxxxxxxxx",
		"78xxxxxxxxxxx",
		"7xxxxxxxxxxxx"
		],
		compared_to_category: "en:whole-milks",
		complete: 0,
		completeness: 0.9,
		correctors_tags: [
		"openfoodfacts-contributors",
		"roboto-app",
		"teolemon",
		"michelpick",
		"darllesonlessa",
		"foodless"
		],
		countries: "en:Brazil",
		countries_hierarchy: [
		"en:brazil"
		],
		countries_lc: "pt",
		countries_tags: [
		"en:brazil"
		],
		created_t: 1577062936,
		creator: "openfoodfacts-contributors",
		data_quality_bugs_tags: [ ],
		data_quality_errors_tags: [ ],
		data_quality_info_tags: [
		"en:packaging-data-incomplete",
		"en:ingredients-percent-analysis-ok",
		"en:ecoscore-extended-data-not-computed",
		"en:food-groups-1-known",
		"en:food-groups-2-known",
		"en:food-groups-3-unknown"
		],
		data_quality_tags: [
		"en:packaging-data-incomplete",
		"en:ingredients-percent-analysis-ok",
		"en:ecoscore-extended-data-not-computed",
		"en:food-groups-1-known",
		"en:food-groups-2-known",
		"en:food-groups-3-unknown",
		"en:ingredients-unknown-score-above-0",
		"en:ingredients-70-percent-unknown",
		"en:ingredients-ingredient-tag-length-greater-than-50",
		"en:ingredients-ingredient-tag-length-greater-than-100",
		"en:nutrition-value-very-high-for-category-energy",
		"en:nutrition-value-very-high-for-category-fat",
		"en:nutrition-value-very-high-for-category-saturated-fat",
		"en:ecoscore-origins-of-ingredients-origins-are-100-percent-unknown",
		"en:ecoscore-production-system-no-label"
		],
		data_quality_warnings_tags: [
		"en:ingredients-unknown-score-above-0",
		"en:ingredients-70-percent-unknown",
		"en:ingredients-ingredient-tag-length-greater-than-50",
		"en:ingredients-ingredient-tag-length-greater-than-100",
		"en:nutrition-value-very-high-for-category-energy",
		"en:nutrition-value-very-high-for-category-fat",
		"en:nutrition-value-very-high-for-category-saturated-fat",
		"en:ecoscore-origins-of-ingredients-origins-are-100-percent-unknown",
		"en:ecoscore-production-system-no-label"
		],
		data_sources: "App - off, Apps, App - smoothie-openfoodfacts",
		data_sources_tags: [
		"app-off",
		"apps",
		"app-smoothie-openfoodfacts"
		],
		ecoscore_data: {
		adjustments: {
		origins_of_ingredients: {
		aggregated_origins: [
		{
		epi_score: "0",
		origin: "en:unknown",
		percent: 100,
		transportation_score: null
		}
		],
		epi_score: 0,
		epi_value: -5,
		origins_from_categories: [
		"en:unknown"
		],
		origins_from_origins_field: [
		"en:unknown"
		],
		transportation_score: 0,
		transportation_scores: {
		ad: 0,
		al: 0,
		at: 0,
		ax: 0,
		ba: 0,
		be: 0,
		bg: 0,
		ch: 0,
		cy: 0,
		cz: 0,
		de: 0,
		dk: 0,
		dz: 0,
		ee: 0,
		eg: 0,
		es: 0,
		fi: 0,
		fo: 0,
		fr: 0,
		gg: 0,
		gi: 0,
		gr: 0,
		hr: 0,
		hu: 0,
		ie: 0,
		il: 0,
		im: 0,
		is: 0,
		it: 0,
		je: 0,
		lb: 0,
		li: 0,
		lt: 0,
		lu: 0,
		lv: 0,
		ly: 0,
		ma: 0,
		mc: 0,
		md: 0,
		me: 0,
		mk: 0,
		mt: 0,
		nl: 0,
		no: 0,
		pl: 0,
		ps: 0,
		pt: 0,
		ro: 0,
		rs: 0,
		se: 0,
		si: 0,
		sj: 0,
		sk: 0,
		sm: 0,
		sy: 0,
		tn: 0,
		tr: 0,
		ua: 0,
		uk: 0,
		us: 0,
		va: 0,
		world: 0,
		xk: 0
		},
		transportation_value: 0,
		transportation_values: {
		ad: 0,
		al: 0,
		at: 0,
		ax: 0,
		ba: 0,
		be: 0,
		bg: 0,
		ch: 0,
		cy: 0,
		cz: 0,
		de: 0,
		dk: 0,
		dz: 0,
		ee: 0,
		eg: 0,
		es: 0,
		fi: 0,
		fo: 0,
		fr: 0,
		gg: 0,
		gi: 0,
		gr: 0,
		hr: 0,
		hu: 0,
		ie: 0,
		il: 0,
		im: 0,
		is: 0,
		it: 0,
		je: 0,
		lb: 0,
		li: 0,
		lt: 0,
		lu: 0,
		lv: 0,
		ly: 0,
		ma: 0,
		mc: 0,
		md: 0,
		me: 0,
		mk: 0,
		mt: 0,
		nl: 0,
		no: 0,
		pl: 0,
		ps: 0,
		pt: 0,
		ro: 0,
		rs: 0,
		se: 0,
		si: 0,
		sj: 0,
		sk: 0,
		sm: 0,
		sy: 0,
		tn: 0,
		tr: 0,
		ua: 0,
		uk: 0,
		us: 0,
		va: 0,
		world: 0,
		xk: 0
		},
		value: -5,
		values: {
		ad: -5,
		al: -5,
		at: -5,
		ax: -5,
		ba: -5,
		be: -5,
		bg: -5,
		ch: -5,
		cy: -5,
		cz: -5,
		de: -5,
		dk: -5,
		dz: -5,
		ee: -5,
		eg: -5,
		es: -5,
		fi: -5,
		fo: -5,
		fr: -5,
		gg: -5,
		gi: -5,
		gr: -5,
		hr: -5,
		hu: -5,
		ie: -5,
		il: -5,
		im: -5,
		is: -5,
		it: -5,
		je: -5,
		lb: -5,
		li: -5,
		lt: -5,
		lu: -5,
		lv: -5,
		ly: -5,
		ma: -5,
		mc: -5,
		md: -5,
		me: -5,
		mk: -5,
		mt: -5,
		nl: -5,
		no: -5,
		pl: -5,
		ps: -5,
		pt: -5,
		ro: -5,
		rs: -5,
		se: -5,
		si: -5,
		sj: -5,
		sk: -5,
		sm: -5,
		sy: -5,
		tn: -5,
		tr: -5,
		ua: -5,
		uk: -5,
		us: -5,
		va: -5,
		world: -5,
		xk: -5
		},
		warning: "origins_are_100_percent_unknown"
		},
		packaging: {
		non_recyclable_and_non_biodegradable_materials: 0,
		packagings: [
		{
		ecoscore_material_score: 62,
		ecoscore_shape_ratio: 1,
		material: "en:tetra-pak",
		number_of_units: 1,
		quantity_per_unit: "200 g",
		quantity_per_unit_unit: "g",
		quantity_per_unit_value: 200,
		recycling: "en:recycle",
		shape: "en:box"
		}
		],
		score: 62,
		value: -4
		},
		production_system: {
		labels: [ ],
		value: 0,
		warning: "no_label"
		},
		threatened_species: { }
		},
		agribalyse: {
		agribalyse_proxy_food_code: "19024",
		co2_agriculture: 1.1070005,
		co2_consumption: 0,
		co2_distribution: 0.015701413,
		co2_packaging: 0.14806368,
		co2_processing: 0.01123937,
		co2_total: 1.486686273,
		co2_transportation: 0.20468131,
		code: "19024",
		dqr: "2.03",
		ef_agriculture: 0.10185498,
		ef_consumption: 0,
		ef_distribution: 0.0045893771,
		ef_packaging: 0.014222522,
		ef_processing: 0.00143687,
		ef_total: 0.1379057301,
		ef_transportation: 0.015801981,
		is_beverage: 1,
		name_en: "Milk, whole, pasteurised",
		name_fr: "Lait entier, pasteurisé",
		score: 53,
		version: "3.1.1"
		},
		grade: "c",
		grades: {
		ad: "c",
		al: "c",
		at: "c",
		ax: "c",
		ba: "c",
		be: "c",
		bg: "c",
		ch: "c",
		cy: "c",
		cz: "c",
		de: "c",
		dk: "c",
		dz: "c",
		ee: "c",
		eg: "c",
		es: "c",
		fi: "c",
		fo: "c",
		fr: "c",
		gg: "c",
		gi: "c",
		gr: "c",
		hr: "c",
		hu: "c",
		ie: "c",
		il: "c",
		im: "c",
		is: "c",
		it: "c",
		je: "c",
		lb: "c",
		li: "c",
		lt: "c",
		lu: "c",
		lv: "c",
		ly: "c",
		ma: "c",
		mc: "c",
		md: "c",
		me: "c",
		mk: "c",
		mt: "c",
		nl: "c",
		no: "c",
		pl: "c",
		ps: "c",
		pt: "c",
		ro: "c",
		rs: "c",
		se: "c",
		si: "c",
		sj: "c",
		sk: "c",
		sm: "c",
		sy: "c",
		tn: "c",
		tr: "c",
		ua: "c",
		uk: "c",
		us: "c",
		va: "c",
		world: "c",
		xk: "c"
		},
		missing: {
		labels: 1,
		origins: 1
		},
		missing_data_warning: 1,
		previous_data: {
		agribalyse: {
		warning: "missing_agribalyse_match"
		},
		grade: null,
		score: null
		},
		score: 44,
		scores: {
		ad: 44,
		al: 44,
		at: 44,
		ax: 44,
		ba: 44,
		be: 44,
		bg: 44,
		ch: 44,
		cy: 44,
		cz: 44,
		de: 44,
		dk: 44,
		dz: 44,
		ee: 44,
		eg: 44,
		es: 44,
		fi: 44,
		fo: 44,
		fr: 44,
		gg: 44,
		gi: 44,
		gr: 44,
		hr: 44,
		hu: 44,
		ie: 44,
		il: 44,
		im: 44,
		is: 44,
		it: 44,
		je: 44,
		lb: 44,
		li: 44,
		lt: 44,
		lu: 44,
		lv: 44,
		ly: 44,
		ma: 44,
		mc: 44,
		md: 44,
		me: 44,
		mk: 44,
		mt: 44,
		nl: 44,
		no: 44,
		pl: 44,
		ps: 44,
		pt: 44,
		ro: 44,
		rs: 44,
		se: 44,
		si: 44,
		sj: 44,
		sk: 44,
		sm: 44,
		sy: 44,
		tn: 44,
		tr: 44,
		ua: 44,
		uk: 44,
		us: 44,
		va: 44,
		world: 44,
		xk: 44
		},
		status: "known"
		},
		ecoscore_grade: "c",
		ecoscore_score: 44,
		ecoscore_tags: [
		"c"
		],
		editors_tags: [
		"darllesonlessa",
		"foodless",
		"michelpick",
		"openfoodfacts-contributors",
		"roboto-app",
		"teolemon"
		],
		emb_codes: "7898215161784,FSC-C014047",
		emb_codes_tags: [
		"7898215161784",
		"fsc-c014047"
		],
		entry_dates_tags: [
		"2019-12-23",
		"2019-12",
		"2019"
		],
		food_groups: "en:milk-and-yogurt",
		food_groups_tags: [
		"en:milk-and-dairy-products",
		"en:milk-and-yogurt"
		],
		id: "7898215151784",
		image_front_small_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/front_pt.20.200.jpg",
		image_front_thumb_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/front_pt.20.100.jpg",
		image_front_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/front_pt.20.400.jpg",
		image_ingredients_small_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/ingredients_pt.22.200.jpg",
		image_ingredients_thumb_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/ingredients_pt.22.100.jpg",
		image_ingredients_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/ingredients_pt.22.400.jpg",
		image_nutrition_small_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/nutrition_pt.24.200.jpg",
		image_nutrition_thumb_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/nutrition_pt.24.100.jpg",
		image_nutrition_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/nutrition_pt.24.400.jpg",
		image_packaging_small_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/packaging_pt.26.200.jpg",
		image_packaging_thumb_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/packaging_pt.26.100.jpg",
		image_packaging_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/packaging_pt.26.400.jpg",
		image_small_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/front_pt.20.200.jpg",
		image_thumb_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/front_pt.20.100.jpg",
		image_url: "https://images.openfoodfacts.org/images/products/789/821/515/1784/front_pt.20.400.jpg",
		images: {
		1: {
		sizes: {
		100: {
		h: 100,
		w: 75
		},
		400: {
		h: 400,
		w: 300
		},
		full: {
		h: 4160,
		w: 3120
		}
		},
		uploaded_t: 1577062937,
		uploader: "openfoodfacts-contributors"
		},
		2: {
		sizes: {
		100: {
		h: 100,
		w: 75
		},
		400: {
		h: 400,
		w: 300
		},
		full: {
		h: 4160,
		w: 3120
		}
		},
		uploaded_t: 1577063285,
		uploader: "openfoodfacts-contributors"
		},
		3: {
		sizes: {
		100: {
		h: 75,
		w: 100
		},
		400: {
		h: 300,
		w: 400
		},
		full: {
		h: 3120,
		w: 4160
		}
		},
		uploaded_t: 1577063296,
		uploader: "openfoodfacts-contributors"
		},
		4: {
		sizes: {
		100: {
		h: 100,
		w: 75
		},
		400: {
		h: 400,
		w: 300
		},
		full: {
		h: 4160,
		w: 3120
		}
		},
		uploaded_t: 1577063348,
		uploader: "openfoodfacts-contributors"
		},
		5: {
		sizes: {
		100: {
		h: 100,
		w: 75
		},
		400: {
		h: 400,
		w: 300
		},
		full: {
		h: 4160,
		w: 3120
		}
		},
		uploaded_t: 1577063414,
		uploader: "openfoodfacts-contributors"
		},
		7: {
		sizes: {
		100: {
		h: 100,
		w: 77
		},
		400: {
		h: 400,
		w: 307
		},
		full: {
		h: 2339,
		w: 1797
		}
		},
		uploaded_t: 1697824539,
		uploader: "michelpick"
		},
		8: {
		sizes: {
		100: {
		h: 51,
		w: 100
		},
		400: {
		h: 205,
		w: 400
		},
		full: {
		h: 703,
		w: 1370
		}
		},
		uploaded_t: 1697824557,
		uploader: "michelpick"
		},
		9: {
		sizes: {
		100: {
		h: 100,
		w: 98
		},
		400: {
		h: 400,
		w: 394
		},
		full: {
		h: 1620,
		w: 1595
		}
		},
		uploaded_t: 1697824576,
		uploader: "michelpick"
		},
		10: {
		sizes: {
		100: {
		h: 64,
		w: 100
		},
		400: {
		h: 255,
		w: 400
		},
		full: {
		h: 854,
		w: 1339
		}
		},
		uploaded_t: 1697824591,
		uploader: "michelpick"
		},
		front_pt: {
		angle: 0,
		coordinates_image_size: "full",
		geometry: "0x0--1--1",
		imgid: "7",
		normalize: null,
		rev: "20",
		sizes: {
		100: {
		h: 100,
		w: 77
		},
		200: {
		h: 200,
		w: 154
		},
		400: {
		h: 400,
		w: 307
		},
		full: {
		h: 2339,
		w: 1797
		}
		},
		white_magic: null,
		x1: "-1",
		x2: "-1",
		y1: "-1",
		y2: "-1"
		},
		ingredients_pt: {
		angle: 0,
		coordinates_image_size: "full",
		geometry: "0x0--1--1",
		imgid: "8",
		normalize: null,
		rev: "22",
		sizes: {
		100: {
		h: 51,
		w: 100
		},
		200: {
		h: 103,
		w: 200
		},
		400: {
		h: 205,
		w: 400
		},
		full: {
		h: 703,
		w: 1370
		}
		},
		white_magic: null,
		x1: "-1",
		x2: "-1",
		y1: "-1",
		y2: "-1"
		},
		nutrition_pt: {
		angle: 0,
		coordinates_image_size: "full",
		geometry: "0x0--1--1",
		imgid: "9",
		normalize: null,
		rev: "24",
		sizes: {
		100: {
		h: 100,
		w: 98
		},
		200: {
		h: 200,
		w: 197
		},
		400: {
		h: 400,
		w: 394
		},
		full: {
		h: 1620,
		w: 1595
		}
		},
		white_magic: null,
		x1: "-1",
		x2: "-1",
		y1: "-1",
		y2: "-1"
		},
		packaging_pt: {
		angle: 0,
		coordinates_image_size: "full",
		geometry: "0x0--1--1",
		imgid: "10",
		normalize: null,
		rev: "26",
		sizes: {
		100: {
		h: 64,
		w: 100
		},
		200: {
		h: 128,
		w: 200
		},
		400: {
		h: 255,
		w: 400
		},
		full: {
		h: 854,
		w: 1339
		}
		},
		white_magic: null,
		x1: "-1",
		x2: "-1",
		y1: "-1",
		y2: "-1"
		}
		},
		informers_tags: [
		"openfoodfacts-contributors",
		"teolemon",
		"michelpick"
		],
		ingredients: [
		{
		id: "pt:creme-de-leite-padronizado-a-17-de-gordura",
		percent_estimate: 62.5,
		percent_max: 100,
		percent_min: 25,
		rank: 1,
		text: "Creme de leite padronizado a 17% de gordura"
		},
		{
		ciqual_proxy_food_code: "19051",
		id: "en:skimmed-milk",
		percent_estimate: 18.75,
		percent_max: 50,
		percent_min: 0,
		processing: "en:powder",
		rank: 2,
		text: "leite desnatado",
		vegan: "no",
		vegetarian: "yes"
		},
		{
		id: "pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		percent_estimate: 9.375,
		percent_max: 33.3333333333333,
		percent_min: 0,
		rank: 3,
		text: "espessantes carragena carboximetilcelulose sódica e alginato de sódio e estabilizantes celulose microcristalina e citrato de sódio"
		},
		{
		has_sub_ingredients: "yes",
		id: "pt:alergicos",
		percent_estimate: 9.375,
		percent_max: 25,
		percent_min: 0,
		rank: 4,
		text: "_ALÉRGICOS"
		},
		{
		id: "pt:contem-leite-e-derivados",
		percent_estimate: 4.6875,
		percent_max: 25,
		percent_min: 0,
		text: "CONTÉM LEITE E DERIVADOS"
		},
		{
		id: "en:lactose",
		percent_estimate: 2.34375,
		percent_max: 12.5,
		percent_min: 0,
		text: "CONTÉM LACTOSE",
		vegan: "no",
		vegetarian: "yes"
		},
		{
		id: "pt:nao-contem-gluten",
		percent_estimate: 2.34375,
		percent_max: 8.33333333333333,
		percent_min: 0,
		text: "NÃO CONTÉM GLÚTEN._"
		}
		],
		ingredients_analysis: {
		en:non-vegan: [
		"en:skimmed-milk",
		"en:lactose"
		],
		en:palm-oil-content-unknown: [
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"pt:nao-contem-gluten"
		],
		en:vegan-status-unknown: [
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"pt:nao-contem-gluten"
		],
		en:vegetarian-status-unknown: [
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"pt:nao-contem-gluten"
		]
		},
		ingredients_analysis_tags: [
		"en:palm-oil-content-unknown",
		"en:non-vegan",
		"en:vegetarian-status-unknown"
		],
		ingredients_from_or_that_may_be_from_palm_oil_n: 0,
		ingredients_from_palm_oil_n: 0,
		ingredients_from_palm_oil_tags: [ ],
		ingredients_hierarchy: [
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"en:skimmed-milk",
		"en:dairy",
		"en:milk",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"en:lactose",
		"pt:nao-contem-gluten"
		],
		ingredients_lc: "pt",
		ingredients_n: 7,
		ingredients_n_tags: [
		"7",
		"1-10"
		],
		ingredients_non_nutritive_sweeteners_n: 0,
		ingredients_original_tags: [
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"en:skimmed-milk",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"en:lactose",
		"pt:nao-contem-gluten"
		],
		ingredients_percent_analysis: 1,
		ingredients_sweeteners_n: 0,
		ingredients_tags: [
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"en:skimmed-milk",
		"en:dairy",
		"en:milk",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"en:lactose",
		"pt:nao-contem-gluten"
		],
		ingredients_text: "Creme de leite padronizado a 17% de gordura, leite em pó desnatado, espessantes carragena carboximetilcelulose sódica e alginato de sódio e estabilizantes celulose microcristalina e citrato de sódio. _ALÉRGICOS: CONTÉM LEITE E DERIVADOS. CONTÉM LACTOSE. NÃO CONTÉM GLÚTEN._",
		ingredients_text_pt: "Creme de leite padronizado a 17% de gordura, leite em pó desnatado, espessantes carragena carboximetilcelulose sódica e alginato de sódio e estabilizantes celulose microcristalina e citrato de sódio. _ALÉRGICOS: CONTÉM LEITE E DERIVADOS. CONTÉM LACTOSE. NÃO CONTÉM GLÚTEN._",
		ingredients_text_with_allergens: "Creme de leite padronizado a 17% de gordura, leite em pó desnatado, espessantes carragena carboximetilcelulose sódica e alginato de sódio e estabilizantes celulose microcristalina e citrato de sódio. <span class="allergen">ALÉRGICOS: <span class="allergen">CONTÉM LEITE E DERIVADOS</span>. <span class="allergen">CONTÉM LACTOSE</span>. NÃO CONTÉM GLÚTEN.</span>",
		ingredients_text_with_allergens_pt: "Creme de leite padronizado a 17% de gordura, leite em pó desnatado, espessantes carragena carboximetilcelulose sódica e alginato de sódio e estabilizantes celulose microcristalina e citrato de sódio. <span class="allergen">ALÉRGICOS: <span class="allergen">CONTÉM LEITE E DERIVADOS</span>. <span class="allergen">CONTÉM LACTOSE</span>. NÃO CONTÉM GLÚTEN.</span>",
		ingredients_that_may_be_from_palm_oil_n: 0,
		ingredients_that_may_be_from_palm_oil_tags: [ ],
		ingredients_with_specified_percent_n: 0,
		ingredients_with_specified_percent_sum: 0,
		ingredients_with_unspecified_percent_n: 6,
		ingredients_with_unspecified_percent_sum: 100,
		ingredients_without_ciqual_codes: [
		"en:lactose",
		"pt:alergicos",
		"pt:contem-leite-e-derivados",
		"pt:creme-de-leite-padronizado-a-17-de-gordura",
		"pt:espessantes-carragena-carboximetilcelulose-sodica-e-alginato-de-sodio-e-estabilizantes-celulose-microcristalina-e-citrato-de-sodio",
		"pt:nao-contem-gluten"
		],
		ingredients_without_ciqual_codes_n: 6,
		interface_version_created: "20120622",
		interface_version_modified: "20150316.jqm2",
		known_ingredients_n: 4,
		labels: "FSC",
		labels_hierarchy: [
		"en:fsc"
		],
		labels_lc: "pt",
		labels_old: "en:FSC",
		labels_tags: [
		"en:fsc"
		],
		lang: "pt",
		languages: {
		en:portuguese: 5
		},
		languages_codes: {
		pt: 5
		},
		languages_hierarchy: [
		"en:portuguese"
		],
		languages_tags: [
		"en:portuguese",
		"en:1"
		],
		last_edit_dates_tags: [
		"2024-04-19",
		"2024-04",
		"2024"
		],
		last_editor: "roboto-app",
		last_image_dates_tags: [
		"2023-10-20",
		"2023-10",
		"2023"
		],
		last_image_t: 1697824591,
		last_modified_by: "roboto-app",
		last_modified_t: 1713546616,
		last_updated_t: 1713546616,
		lc: "pt",
		link: "www.piracanjuba.com.br",
		main_countries_tags: [ ],
		manufacturing_places: "Bela Vista de Goiás",
		manufacturing_places_tags: [
		"bela-vista-de-goias"
		],
		max_imgid: "10",
		minerals_tags: [ ],
		misc_tags: [
		"en:nutriscore-computed",
		"en:nutrition-fruits-vegetables-nuts-estimate-from-ingredients",
		"en:nutrition-all-nutriscore-values-known",
		"en:nutrition-fruits-vegetables-legumes-estimate-from-ingredients",
		"en:nutriscore-2021-different-from-2023",
		"en:nutriscore-2021-better-than-2023",
		"en:nutriscore-2021-d-2023-e",
		"en:packagings-number-of-components-1",
		"en:packagings-not-complete",
		"en:packagings-not-empty-but-not-complete",
		"en:packagings-not-empty",
		"en:ecoscore-extended-data-not-computed",
		"en:ecoscore-missing-data-warning",
		"en:ecoscore-missing-data-labels",
		"en:ecoscore-missing-data-origins",
		"en:ecoscore-computed",
		"en:ecoscore-changed",
		"en:ecoscore-grade-changed"
		],
		no_nutrition_data: "",
		nova_group: 4,
		nova_group_debug: "",
		nova_groups: "4",
		nova_groups_markers: {
		4: [
		[
		"ingredients",
		"en:lactose"
		]
		]
		},
		nova_groups_tags: [
		"en:4-ultra-processed-food-and-drink-products"
		],
		nucleotides_tags: [ ],
		nutrient_levels: {
		fat: "moderate",
		salt: "low",
		saturated-fat: "high",
		sugars: "low"
		},
		nutrient_levels_tags: [
		"en:fat-in-moderate-quantity",
		"en:saturated-fat-in-high-quantity",
		"en:sugars-in-low-quantity",
		"en:salt-in-low-quantity"
		],
		nutriments: {
		alcohol: 0,
		alcohol_100g: 0,
		alcohol_serving: 0,
		alcohol_unit: "% vol",
		alcohol_value: 0,
		carbohydrates: 4.3,
		carbohydrates_100g: 4.3,
		carbohydrates_serving: 0.645,
		carbohydrates_unit: "g",
		carbohydrates_value: 4.3,
		energy: 761,
		energy-kcal: 182,
		energy-kcal_100g: 182,
		energy-kcal_serving: 27.3,
		energy-kcal_unit: "kcal",
		energy-kcal_value: 182,
		energy-kcal_value_computed: 181.8,
		energy_100g: 761,
		energy_serving: 114,
		energy_unit: "kcal",
		energy_value: 182,
		fat: 17,
		fat_100g: 17,
		fat_serving: 2.55,
		fat_unit: "g",
		fat_value: 17,
		fiber: 0,
		fiber_100g: 0,
		fiber_serving: 0,
		fiber_unit: "g",
		fiber_value: 0,
		fruits-vegetables-legumes-estimate-from-ingredients_100g: 0,
		fruits-vegetables-legumes-estimate-from-ingredients_serving: 0,
		fruits-vegetables-nuts-estimate-from-ingredients_100g: 0,
		fruits-vegetables-nuts-estimate-from-ingredients_serving: 0,
		nova-group: 4,
		nova-group_100g: 4,
		nova-group_serving: 4,
		nutrition-score-fr: 12,
		nutrition-score-fr_100g: 12,
		proteins: 2.9,
		proteins_100g: 2.9,
		proteins_serving: 0.435,
		proteins_unit: "g",
		proteins_value: 2.9,
		salt: 0.08,
		salt_100g: 0.08,
		salt_serving: 0.012,
		salt_unit: "g",
		salt_value: 0.08,
		saturated-fat: 11,
		saturated-fat_100g: 11,
		saturated-fat_serving: 1.65,
		saturated-fat_unit: "g",
		saturated-fat_value: 11,
		sodium: 0.032,
		sodium_100g: 0.032,
		sodium_serving: 0.0048,
		sodium_unit: "g",
		sodium_value: 0.032,
		sugars: 4.3,
		sugars_100g: 4.3,
		sugars_serving: 0.645,
		sugars_unit: "g",
		sugars_value: 4.3
		},
		nutriscore: {
		2021: {
		category_available: 1,
		data: {
		energy: 761,
		energy_points: 2,
		energy_value: 761,
		fiber: 0,
		fiber_points: 0,
		fiber_value: 0,
		fruits_vegetables_nuts_colza_walnut_olive_oils: 0,
		fruits_vegetables_nuts_colza_walnut_olive_oils_points: 0,
		fruits_vegetables_nuts_colza_walnut_olive_oils_value: 0,
		is_beverage: 0,
		is_cheese: 0,
		is_fat: 0,
		is_water: 0,
		negative_points: 12,
		positive_points: 0,
		proteins: 2.9,
		proteins_points: 1,
		proteins_value: 2.9,
		saturated_fat: 11,
		saturated_fat_points: 10,
		saturated_fat_value: 11,
		sodium: 32,
		sodium_points: 0,
		sodium_value: 32,
		sugars: 4.3,
		sugars_points: 0,
		sugars_value: 4.3
		},
		grade: "d",
		nutrients_available: 1,
		nutriscore_applicable: 1,
		nutriscore_computed: 1,
		score: 12
		},
		2023: {
		category_available: 1,
		data: {
		components: {
		negative: [
		{
		id: "energy_from_saturated_fat",
		points: 3,
		points_max: 10,
		unit: "kJ",
		value: 407
		},
		{
		id: "sugars",
		points: 3,
		points_max: 10,
		unit: "g",
		value: 4.3
		},
		{
		id: "saturated_fat_ratio",
		points: 10,
		points_max: 10,
		unit: "%",
		value: 64.7
		},
		{
		id: "salt",
		points: 0,
		points_max: 20,
		unit: "g",
		value: 0.08
		},
		{
		id: "non_nutritive_sweeteners",
		points: 0,
		points_max: 4,
		unit: "number",
		value: 0
		}
		],
		positive: [
		{
		id: "proteins",
		points: 6,
		points_max: 7,
		unit: "g",
		value: 2.9
		},
		{
		id: "fiber",
		points: 0,
		points_max: 5,
		unit: "g",
		value: 0
		},
		{
		id: "fruits_vegetables_legumes",
		points: 0,
		points_max: 6,
		unit: "%",
		value: 0
		}
		]
		},
		count_proteins: 1,
		count_proteins_reason: "beverage",
		is_beverage: 1,
		is_cheese: 0,
		is_fat_oil_nuts_seeds: 1,
		is_red_meat_product: 0,
		is_water: 0,
		negative_points: 16,
		negative_points_max: 54,
		positive_nutrients: [
		"proteins",
		"fiber",
		"fruits_vegetables_legumes"
		],
		positive_points: 6,
		positive_points_max: 18
		},
		grade: "e",
		nutrients_available: 1,
		nutriscore_applicable: 1,
		nutriscore_computed: 1,
		score: 10
		}
		},
		nutriscore_2021_tags: [
		"d"
		],
		nutriscore_2023_tags: [
		"e"
		],
		nutriscore_data: {
		energy: 761,
		energy_points: 2,
		energy_value: 761,
		fiber: 0,
		fiber_points: 0,
		fiber_value: 0,
		fruits_vegetables_nuts_colza_walnut_olive_oils: 0,
		fruits_vegetables_nuts_colza_walnut_olive_oils_points: 0,
		fruits_vegetables_nuts_colza_walnut_olive_oils_value: 0,
		grade: "d",
		is_beverage: 0,
		is_cheese: 0,
		is_fat: 0,
		is_water: 0,
		negative_points: 12,
		positive_points: 0,
		proteins: 2.9,
		proteins_points: 1,
		proteins_value: 2.9,
		saturated_fat: 11,
		saturated_fat_points: 10,
		saturated_fat_value: 11,
		score: 12,
		sodium: 32,
		sodium_points: 0,
		sodium_value: 32,
		sugars: 4.3,
		sugars_points: 0,
		sugars_value: 4.3
		},
		nutriscore_grade: "d",
		nutriscore_score: 12,
		nutriscore_score_opposite: -12,
		nutriscore_tags: [
		"d"
		],
		nutriscore_version: "2021",
		nutrition_data: "on",
		nutrition_data_per: "100g",
		nutrition_data_prepared_per: "100g",
		nutrition_grade_fr: "d",
		nutrition_grades: "d",
		nutrition_grades_tags: [
		"d"
		],
		nutrition_score_beverage: 0,
		nutrition_score_debug: "",
		nutrition_score_warning_fruits_vegetables_legumes_estimate_from_ingredients: 1,
		nutrition_score_warning_fruits_vegetables_legumes_estimate_from_ingredients_value: 0,
		nutrition_score_warning_fruits_vegetables_nuts_estimate_from_ingredients: 1,
		nutrition_score_warning_fruits_vegetables_nuts_estimate_from_ingredients_value: 0,
		other_nutritional_substances_tags: [ ],
		packaging: "Caixinha",
		packaging_hierarchy: [
		"pt:Caixinha"
		],
		packaging_lc: "pt",
		packaging_materials_tags: [
		"en:tetra-pak"
		],
		packaging_old: "caixinha",
		packaging_recycling_tags: [
		"en:recycle"
		],
		packaging_shapes_tags: [
		"en:box"
		],
		packaging_tags: [
		"pt:caixinha"
		],
		packagings: [
		{
		material: "en:tetra-pak",
		number_of_units: 1,
		quantity_per_unit: "200 g",
		quantity_per_unit_unit: "g",
		quantity_per_unit_value: 200,
		recycling: "en:recycle",
		shape: "en:box"
		}
		],
		packagings_materials: {
		all: { },
		en:unknown: { }
		},
		packagings_n: 1,
		photographers_tags: [
		"openfoodfacts-contributors",
		"michelpick"
		],
		pnns_groups_1: "Milk and dairy products",
		pnns_groups_1_tags: [
		"milk-and-dairy-products",
		"known"
		],
		pnns_groups_2: "Milk and yogurt",
		pnns_groups_2_tags: [
		"milk-and-yogurt",
		"known"
		],
		popularity_key: 22900000013,
		popularity_tags: [
		"bottom-25-percent-scans-2019",
		"bottom-20-percent-scans-2019",
		"top-85-percent-scans-2019",
		"top-90-percent-scans-2019",
		"top-50-br-scans-2019",
		"top-100-br-scans-2019",
		"top-500-br-scans-2019",
		"top-1000-br-scans-2019",
		"top-5000-br-scans-2019",
		"top-10000-br-scans-2019",
		"top-50000-br-scans-2019",
		"top-100000-br-scans-2019",
		"top-country-br-scans-2019",
		"top-100000-scans-2020",
		"at-least-5-scans-2020",
		"top-75-percent-scans-2020",
		"top-80-percent-scans-2020",
		"top-85-percent-scans-2020",
		"top-90-percent-scans-2020",
		"top-50-br-scans-2020",
		"top-100-br-scans-2020",
		"top-500-br-scans-2020",
		"top-1000-br-scans-2020",
		"top-5000-br-scans-2020",
		"top-10000-br-scans-2020",
		"top-50000-br-scans-2020",
		"top-100000-br-scans-2020",
		"top-country-br-scans-2020",
		"at-least-5-br-scans-2020",
		"top-100000-scans-2021",
		"at-least-5-scans-2021",
		"top-75-percent-scans-2021",
		"top-80-percent-scans-2021",
		"top-85-percent-scans-2021",
		"top-90-percent-scans-2021",
		"top-50-br-scans-2021",
		"top-100-br-scans-2021",
		"top-500-br-scans-2021",
		"top-1000-br-scans-2021",
		"top-5000-br-scans-2021",
		"top-10000-br-scans-2021",
		"top-50000-br-scans-2021",
		"top-100000-br-scans-2021",
		"top-country-br-scans-2021",
		"top-75-percent-scans-2022",
		"top-80-percent-scans-2022",
		"top-85-percent-scans-2022",
		"top-90-percent-scans-2022",
		"top-100-br-scans-2022",
		"top-500-br-scans-2022",
		"top-1000-br-scans-2022",
		"top-5000-br-scans-2022",
		"top-10000-br-scans-2022",
		"top-50000-br-scans-2022",
		"top-100000-br-scans-2022",
		"top-country-br-scans-2022",
		"top-100000-scans-2023",
		"at-least-5-scans-2023",
		"top-75-percent-scans-2023",
		"top-80-percent-scans-2023",
		"top-85-percent-scans-2023",
		"top-90-percent-scans-2023",
		"top-50-br-scans-2023",
		"top-100-br-scans-2023",
		"top-500-br-scans-2023",
		"top-1000-br-scans-2023",
		"top-5000-br-scans-2023",
		"top-10000-br-scans-2023",
		"top-50000-br-scans-2023",
		"top-100000-br-scans-2023",
		"top-country-br-scans-2023",
		"at-least-5-br-scans-2023"
		],
		product_name: "Creme de Leite Piracanjuba",
		product_name_pt: "Creme de Leite Piracanjuba",
		product_quantity: "200",
		product_quantity_unit: "g",
		purchase_places: "Brasil",
		purchase_places_tags: [
		"brasil"
		],
		quantity: "200 g",
		removed_countries_tags: [ ],
		rev: 35,
		scans_n: 7,
		selected_images: {},
		serving_quantity: "15",
		serving_quantity_unit: "g",
		serving_size: "15g",
		sortkey: 1588956320,
		states: "en:to-be-completed, en:nutrition-facts-completed, en:ingredients-completed, en:expiration-date-to-be-completed, en:packaging-code-completed, en:characteristics-to-be-completed, en:origins-to-be-completed, en:categories-completed, en:brands-completed, en:packaging-completed, en:quantity-completed, en:product-name-completed, en:photos-validated, en:packaging-photo-selected, en:nutrition-photo-selected, en:ingredients-photo-selected, en:front-photo-selected, en:photos-uploaded",
		states_hierarchy: [],
		states_tags: [],
		stores: "Atacadao",
		stores_tags: [
		"atacadao"
		],
		teams: "chocolatine,la-robe-est-bleue",
		teams_tags: [
		"chocolatine",
		"la-robe-est-bleue"
		],
		traces: "",
		traces_from_ingredients: "",
		traces_from_user: "(pt) ",
		traces_hierarchy: [ ],
		traces_tags: [ ],
		unique_scans_n: 7,
		unknown_ingredients_n: 5,
		unknown_nutrients_tags: [ ],
		update_key: "20240209",
		vitamins_tags: [ ],
		weighers_tags: [ ]
		},
		status: 1,
		status_verbose: "product found"
		}`
}
