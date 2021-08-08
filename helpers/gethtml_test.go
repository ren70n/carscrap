package helpers

import (
	"testing"
	"fmt"
)

var domains = [3]string{
	"https://google.com",
	"https://example.com",
	"https://golang.org",
}
var fail_domains = [4]string{
	"htsp",
	"ftp:/zyz",
	"http://+_",
	"https://golang.or",
}

func TestGetHtmlValid(t *testing.T){
	for _,v := range domains{
		html,err := GetHtml(v)

		if err != nil {
			t.Fatalf("%q resulted in error: %v", v, err)
		}
		if len(html)<1{
			t.Fatalf("%q returned 0 length", v)	
		}
	}
}

func TestGetHtmlInvalid(t *testing.T){
	for _, v := range fail_domains{
		fmt.Println(v)
		html,err := GetHtml(v)
		if err == nil{
			t.Fatalf(html)
		}
	}
}