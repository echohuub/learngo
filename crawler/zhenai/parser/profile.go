package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// <div class="des f-cl" data-v-3c42fade>阿坝 | 50岁 | 高中及以下 | 离异 | 158cm | 3000元以下</div>
var contentRe = regexp.MustCompile(`class="des f-cl"[^>]*>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	matches := contentRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		content := parseContent(string(m[1]))
		profile.Address = content[0]
		profile.Education = content[2]
		profile.Marriage = content[3]
		profile.Income = content[5]
		age, err := strconv.Atoi(content[1][0 : utf8.RuneCountInString(content[1])-1])
		if err == nil {
			profile.Age = age
		}
		height, err := strconv.Atoi(content[4][0 : len(content[4])-2])
		if err == nil {
			profile.Height = height
		}
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func parseContent(content string) []string {
	result := []string{}
	list := strings.Split(content, "|")
	for _, item := range list {
		result = append(result, strings.TrimSpace(item))
	}
	return result
}
