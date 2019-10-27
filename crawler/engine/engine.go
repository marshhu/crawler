package engine

import(
	"log"
	"crawler/crawler/fetcher"
)

//Run func
func Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body,err := fetcher.Fetch(r.URL)//抓取数据
		log.Printf("Fetching %s",r.URL)
		if err != nil{
			log.Printf("Fetcher:error "+"fetching url %s: %v",r.URL,err)
			continue
		}

		parseResult := r.ParserFunc(body)//解析数据

		requests = append(requests,parseResult.Requests...)

		for _,v := range parseResult.Items{
			log.Printf("got item %v",v)
		}
	}
}