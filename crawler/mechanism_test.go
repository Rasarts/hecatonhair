package crawler

import (
	"testing"
	"time"
)

func TestCrawlerCanGetDocymentByConfig(test *testing.T) {
	smartphonesPage := Page{
		Path: "smartfony-i-svyaz/smartfony-205",
		PageInPaginationSelector: ".pagination-list .pagination-item",
		PageParamPath:            "/f/page=",
		ItemConfig: ItemConfig{
			ItemSelector:        ".grid-view .product-tile",
			NameOfItemSelector:  ".product-tile-title",
			PriceOfItemSelector: ".product-price-current",
		},
	}

	configuration := EntityConfig{
		Company: Company{
			Iri:        "http://www.mvideo.ru/",
			Name:       "M.Video",
			Categories: []string{"Телефоны"},
		},
		Pages: []Page{smartphonesPage},
	}

	mechanism := NewCrawler()

	go mechanism.RunWithConfiguration(configuration)

	isRightItems := false

	go func() {
		time.Sleep(time.Second * 3)
		close(mechanism.Items)
	}()

	for item := range mechanism.Items {
		if item.Name != "" && item.Price != "" {
			isRightItems = true
			break
		}
	}

	if isRightItems == false {
		test.Fail()
	}
}
