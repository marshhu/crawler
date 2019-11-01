package parser

import(
	"regexp"
	"crawler/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`

//ParseCity parse city info
func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	match := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range match{
		result.Items = append(result.Items,"User "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			URL:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return result
}