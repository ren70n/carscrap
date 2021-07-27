package adapters

import (
	"github.com/ren70n/carscrap/helpers"
	"fmt"
)

type CarArray struct{
	Brand 		string
	Model		[]string
	Year 		string
	Link 		string
	Source 		string
	Image 		string
	Engine 		string	// size in dl
	Fueltype 	string	// diesel/petrol/LPG
	Power 		string	// in HP
	Mileage 	string	// in km
	Price		float32
	Currency	helpers.Currency
}

func (ca CarArray)Print(){
	fmt.Println("=======================")
	fmt.Println("manufacturer: ",ca.Brand)
	fmt.Println("model: ",ca.Model)
	fmt.Println("price: ", ca.Price)
	fmt.Println("found via: ",ca.Source)
	fmt.Println("link: ",ca.Link)
}