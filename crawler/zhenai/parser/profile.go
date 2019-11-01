package parser

import(
	"strings"
	"regexp"
	"crawler/crawler/engine"
	"crawler/crawler/model"
)

var  baseInfoRe = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^<]+)</div>`)
var imageRe = regexp.MustCompile(`background-image:url(([^?]+)?[^)]*)`)

//ParseProfile parse city info
func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}

	match := baseInfoRe.FindSubmatch(contents,-1)
    if match != nil{
		baseInfo := string(match[1]);
		infos := strings.Split(baseInfo,"|")
		if len(infos) >= 6{
			profile.Address = strings.Trim(infos[0],"")
			profile.Age = strings.Trim(infos[1],"")
			profile.Education = strings.Trim(infos[2],"")
			profile.Marriage = strings.Trim(infos[3],"")
			profile.Height = strings.Trim(infos[4],"")
			profile.Income = strings.Trim(infos[5],"")
		}
	}

	match = imageRe.FindAllSubmatch(contents,-1)
	if match != null{
		for _,m := range match{
			profile.Images = append(profile.Images,string(m[1]))
		}
	}

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}

	return result
}