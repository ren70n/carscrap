package adapters

import (
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
)
var url = "https://www.autohero.com/"

func AutoheroGetter() []CarArray{
	// build query
	query := url+"pl/search/?sort=PUBLISHED_AT_DESC"//&priceMax=35000&bodyType=STATION_WAGON&bodyType=SUV&bodyType=VAN_TRANSPORTER"
	
	// get data
	resp,err := http.Get(query)
	defer resp.Body.Close()

	if err != nil{
		log.Fatal(err)
	}

	html, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
	}

	// as a return I have a list of divs
	divToParse := clearHtml(string(html))

	//extracting list possitions
	divs := make([]string,0)

	for len(divToParse)>0{
		div,ix := extractDiv(divToParse)
		divToParse = divToParse[ix:]

		divs = append(divs,div)
	}

	// parse data

	cars := make([]CarArray,0)

	for _,div:=range divs{
		link := getLink(div)
		brand,model := getBrandModel(link)
		image := getImage(div)
		price := getPrice(div)

		if brand == ""{
			continue
		}

		ca := CarArray{
			Link:link,
			Source: "AutoHero",
			Brand: brand,
			Model: model,
			Image: image,
			Price: price,
		}
		ca.Print()
		cars = append(cars,ca)
	}

	return cars
}

func clearHtml(html string)string{

	ix := strings.Index(html,`<div class="ReactVirtualized__Grid__innerScrollContainer`)

	html = html[ix:]

	// here we have a div containing the list
	html,_ = extractDiv(html)


	// I have to remove 1st bracket and last /div

	ix = strings.Index(html,">")
	html = html[ix+1:len(html)-6]

	return html
}


func extractDiv(div string)(string,int){

	if len(div)<5{return "",len(div)}
	i:=0
	do:=0
	for i<len(div)-5{
		if div[i:i+4] == "<div"{
			do++
		}
		if div[i:i+5] == "</div"{
			do--
		}
		if do==0{break}
		i++
	}

	return div[:i+6],i+6

}

func clearDiv(div string)string{
	ib := strings.Index(div,">")
	div = div[ib+1:len(div)-6]

	return div
}

func getLink(div string)string{
	
	io := strings.Index(div,`href="`)
	if io < 0{
		return ""
	}
	ic := strings.Index(div,`/"`)
	// links = append(links,url+div[io+6:ic+1])

	return url+div[io+7:ic+1]
}

// privide full link string for this particular car 
func getBrandModel(link string)(string,[]string){
	sarr := strings.Split(link,"/")

	if len(sarr)<6{return "",nil}

	bm := strings.Split(sarr[5],"-")

	bi := 0

	//ifology as there are brands with 2-part-names
	switch bm[0]{
		case "land": bi = 1
		case "range": bi = 1
		case "alfa": bi = 1
		case "mercedes": bi = 1
	}


	brand := strings.Join(bm[:bi+1]," ")
	model := bm[bi+1:]

	return brand,model
}

// Image is not appearing when grabbing this way
func getImage(div string)string{
	// ix := strings.Index(div,`<img src="https:`)

	// if ix<0{return ""}

	// div = div[ix:]
	// ix = strings.Index(div,`"`)
	// div = div[:ix]

	// log.Println(div)

	return ""
}

func getPrice(div string)float32{
	ix := strings.Index(div,"price")

	if ix<0{return -1.}

	lx := strings.Index(div[ix:],"</div>")

	ps := div[ix:ix+lx]

	ix = strings.Index(ps,">")
	ps = ps[ix+1:len(ps)-5]

	ps = strings.TrimSpace(ps)

	pps :=""
	for _,v:=range ps{
		if v>'0'-1 && v<'9'+1{
			pps+=string(v)
		} 
	}

	price,err := strconv.Atoi(pps)

	if err != nil { return -1. }

	return float32(price)
}