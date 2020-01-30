package parser

import (
	"io/ioutil"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/1051606781", "浅浅")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",
			result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1051606781",
		Type: "zhenai",
		Id:   "1051606781",
		Payload: model.Profile{
			Name:      "浅浅",
			Age:       25,
			Height:    162,
			Income:    "3001-5000元",
			Marriage:  "未婚",
			Education: "大专",
			Address:   "长治",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
