package autohero

import (
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/ren70n/carscrap/adapters"
)
const (
	url = "https://www.autohero.com/"
)

var fType = map[string]adapters.FuelType{
	"Benzyna": adapters.Petrol,
	"Diesel": adapters.Diesel,
	"Elektryczny": adapters.Electric,
	"Hybryda": adapters.Hybrid,
}

var tType = map[string]adapters.TransmissionType{
	"Automatyczna": adapters.Automatic,
	"Manualna": adapters.Manual,
	"Pół-automatyczna": adapters.SemiAutomatic,
	"Dwusprzęgłowa skrzynia biegów": adapters.TwinClutch,
}

func AutoheroGetter() []adapters.CarArray{
	// build query
	query := url+"pl/search"///?sort=PUBLISHED_AT_DESC&priceMax=35000&bodyType=STATION_WAGON&bodyType=SUV&bodyType=VAN_TRANSPORTER"
	
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

	cars := make([]adapters.CarArray,0)

	// for all divs we received we grab data (each div is one car)

	for _,div:=range divs{
		link := getLink(div)
		brand,model := getBrandModel(link)

		// if link is malformed, do not contain the car's name, there is no reason to make the rest of the process
		if brand == ""{
			continue
		}

		// not working yet
		image := getImage(div)

		price := getPrice(div)
		engines,enginet := getEngine(div)

		details := getDetails(div)
		year,_:= strconv.Atoi(details[0])

		available := isAvailable(link)

		ca := adapters.CarArray{
			Link:link,
			Source: "AutoHero",
			Brand: brand,
			Model: model,
			Image: image,
			Price: price,
			EngineSize: engines,
			EngineType: enginet,
			Year: year,
			FuelType: fType[details[1]],
			// must be changed to int
			Mileage: details[2],
			Transmission: tType[details[3]],
			Available: available,

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

	bm := strings.Split(sarr[4],"-")

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

func getEngine(div string)(string,string){
	ix := strings.Index(div,`<h3 class="subtitle___`)
	div = div[ix+1:]

	ix = strings.Index(div,">")
	div = div[ix+1:]

	ix = strings.Index(div,"<")
	div = div[:ix]

	esize:=""

	for _,v := range div{
		if v>='0' && v<='9' || v=='.'{
			esize+=string(v)
		}else{
			break
		}
	}

	return esize,div[len(esize):]
}

func getDetails(div string)[]string{
	ix := strings.Index(div,`spec-list`)

	div = div[ix:]
	ix = strings.Index(div,">")
	div = div[ix+1:]

	ix = strings.Index(div,"</ul>")
	div = div[:ix]


	toret := make([]string,0)
	for len(div)>4{
		ix = strings.Index(div,"-->")+3
		if ix<3{break}

		div = div[ix:]

		sx:=strings.Index(div,"</li>")
		toret = append(toret,div[:sx])
		div = div[sx:]
	}

	return toret
}

func isAvailable(link string)bool{
	
	// get data
	resp,err := http.Get(link)
	defer resp.Body.Close()

	if err != nil{
		log.Fatal(err)
	}

	html, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
	}

	ix := strings.Index(string(html),">Ten samochód jest zarezerwowany<")

	if ix >-1 {return false}

	return true
}