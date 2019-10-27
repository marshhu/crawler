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
	for _,m := range match{
		// 	for _,subMatch := range m{
		// 		fmt.Printf("%s ",subMatch)
		// 	}
		//    fmt.Println()
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			URL:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return result
}