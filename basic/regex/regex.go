package main

import (
	"fmt"
	"regexp"
	"strings"
)

const text = `
My email is qingbao@gmail.com@abc.com
email1 is abc@def.org
email2 is     kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindAllString(text, -1)
	//fmt.Println(match)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}

	fmt.Println(parseData())
	fmt.Println(parseName([]byte(nameContent)))
}

const data = `<div class="des f-cl" data-v-3c42fade>阿坝 | 50岁 | 高中及以下 | 离异 | 158cm | 3000元以下</div>`

func parseData() []string {
	re := regexp.MustCompile(`class="des f-cl"[^>]*>([^<]+)</div>`)
	//fmt.Println(re.FindAllString(data, -1))
	match := re.FindAllStringSubmatch(data, -1)
	result := []string{}
	for _, m := range match {

		fmt.Println(m[1])
		split := strings.Split(string(m[1]), "|")
		for _, item := range split {
			//fmt.Println(strings.TrimSpace(item))
			result = append(result, strings.TrimSpace(item))
		}
	}
	return result
}

const nameContent = `<span class="nickName" data-v-3c42fade>鱼诗诗</span>`

var nameRe = regexp.MustCompile(`class="nickName"[^>]*>([^<]+)</span>`)

func parseName(contents []byte) string {
	match := nameRe.FindSubmatch(contents)
	if match != nil {
		fmt.Println(string(match[1]))
	}
	return ""
}
