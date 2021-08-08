package helpers

import 	(
	"net/http"
	"io/ioutil"
	// "errors"
)

func GetHtml(link string)(string,error){

	// get data
	resp,err := http.Get(link)

	if err != nil{
		return "", err
	}
	// we should not add response to close heap untill we are not sure the response is correct
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		return "", err
	}

	return string(html), nil
}