package parser

import (
	"testing"
	"io/ioutil"
	"crawler/crawler/model"
)

func TestParseProfile(t *testing.T){
	contents,err := ioutil.ReadFile("profile_test_data.html")

	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents)

	if len(result.Items) != 1{
		t.Errorf("Items should contain 1 elemnt;but was %v",result.Items)
	}
	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name : "Angelanancy",
		Gender : "",
		Age : "23岁",
		Height : "168cm",
		Address : "广州",
		Income : "3001-5000元",
		Marriage : "未婚",
		Education : "大学本科",
		Images :[]string{
		"https://photo.zastatic.com/images/photo/429364/1717455853/4579568848494121.png",
		"https://photo.zastatic.com/images/photo/263409/1053635941/5821820458801912.jpg",
		"https://photo.zastatic.com/images/photo/407388/1629549069/4313172023383804.png",
		"https://photo.zastatic.com/images/photo/449074/1796295982/3756984807745179.jpg",
		"https://photo.zastatic.com/images/photo/429364/1717455853/4579568848494121.png",
		},
	}
	if profile != expected{
		t.Errorf("expected %v;but was %v",expected,profile)
	}
}