package parser

import(
	"regexp"
	"crawler/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//ParseCityList parse city info
func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	i := 1;
	for _,m := range match{
		result.Items = append(result.Items,"City "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			URL:string(m[1]),
			ParserFunc:ParseCity,
		})
		i++
		if i == 10{
			break
		}
	}
	return result
}