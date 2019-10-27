package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}

