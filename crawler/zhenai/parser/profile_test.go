package parser

import (
	"golearning/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element. but wa %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expect := model.Profile{
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
	}

	if profile != expect {
		t.Errorf("expect %v; but was %v", expect, profile)
	}
}
