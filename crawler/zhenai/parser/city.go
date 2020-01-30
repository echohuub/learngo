package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

// <a href="http://album.zhenai.com/u/1378032056" target="_blank">鱼诗诗</a>
var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParseFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}

	// 下一页
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
