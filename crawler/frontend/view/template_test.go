package view

import (
	"golearning/crawler/engine"
	"golearning/crawler/frontend/view/model"
	common "golearning/crawler/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	view := CreateSearchResultView("template.html")
	page := model.SearchResult{}
	out, err := os.Create("template_test.html")
	if err != nil {
		panic(err)
	}

	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: common.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)

	if err != nil {
		panic(err)
	}
}
