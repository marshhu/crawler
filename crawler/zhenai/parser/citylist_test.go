package parser

import(
	"testing"
	"io/ioutil"
)
func TestParseCityList(t *testing.T){
	contents,err := ioutil.ReadFile("citylist_test.data.html")
	if err != nil{
		panic(err)
	}
	result := ParseCityList(contents)

	constResult := 470
	if len(result.Requests) != constResult{
		t.Errorf("request shuld have %d requests; but have %d ",constResult, len(result.Requests))
	}

	if len(result.Items) != constResult{
		t.Errorf("request shuld have %d Items; but have %d ",constResult, len(result.Items))
	}
}