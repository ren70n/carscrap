package adapters

import (
	"github.com/ren70n/carscrap/helpers"
	"fmt"
)

type FuelType int
const (
	Petrol FuelType = iota
	Diesel
)

type DriveType int
const (
	Manual DriveType = iota
	Automat
	TwinClutch

)

// var FuelType = map[string]string{
// 	"Benzyna":"petrol",
// 	"Diesel": "diesel",
// }
// var DriveType = map[string]string{
// 	"Benzyna":"petrol",
// 	"Diesel": "diesel",
// }

type CarArray struct{
	Brand 		string
	Model		[]string
	Year 		int
	Link 		string
	Source 		string
	Image 		string
	EngineSize 	string	// size in dl
	EngineType	string	// some fancy additional info about engine
	Drive		string  
	FuelType 	string	// diesel/petrol/LPG
	Power 		string	// in HP
	Mileage 	string	// in km
	Price		float32
	Currency	helpers.Currency
}

func (ca CarArray)Print(){
	fmt.Println("=======================")
	fmt.Println("manufacturer: ",ca.Brand)
	fmt.Println("model: ",ca.Model)
	fmt.Println("engine size:",ca.EngineSize)
	fmt.Println("engine tyoe:",ca.EngineType)
	fmt.Println("fuel:",ca.FuelType)
	fmt.Println("year: ",ca.Year)
	fmt.Println("price: ", ca.Price)
	fmt.Println("found via: ",ca.Source)
	fmt.Println("link: ",ca.Link)
}