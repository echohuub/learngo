package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "浅浅")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",
			result.Items)
	}

	profile := result.Items[0]

	expected := model.Profile{
		Name:      "浅浅",
		Age:       25,
		Height:    162,
		Income:    "3001-5000元",
		Marriage:  "未婚",
		Education: "大专",
		Address:   "长治",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
