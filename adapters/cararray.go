package adapters

import (
	"github.com/ren70n/carscrap/helpers"
	"fmt"
)

type FuelType int
const (
	Petrol FuelType = iota
	Diesel
	Electric
	Hybrid
)

var ftStringEN = map[FuelType]string{
	Petrol: "Petrol",
	Diesel: "Diesel",
	Electric: "Electric",
	Hybrid:	"Hybrid",
}

type TransmissionType int
const (
	Manual TransmissionType = iota
	Automatic
	TwinClutch
	SemiAutomatic
)
var ttString = map[TransmissionType]string {
	Manual: "Manual",
	Automatic: "Automatic",
	TwinClutch: "Twin clutch",
	SemiAutomatic: "Semi-automatic",
}

type CarArray struct{
	Brand 		string
	Model		[]string
	Year 		int
	Link 		string
	Source 		string
	Image 		string
	EngineSize 	string	// size in dl
	EngineType	string	// some fancy additional info about engine
	Transmission	TransmissionType  
	FuelType 	FuelType	// diesel/petrol/LPG
	Power 		string	// in HP
	Mileage 	string	// in km
	Price		float32
	Currency	helpers.Currency
	Available	bool
}

func (ca CarArray)Print(){
	fmt.Println("=======================")
	if !ca.Available{
		fmt.Println("- NOT Available -")	
	}
	fmt.Println("manufacturer: ",ca.Brand)
	fmt.Println("model: ",ca.Model)
	fmt.Println("engine size:",ca.EngineSize)
	fmt.Println("engine tyoe:",ca.EngineType)
	fmt.Println("transmission: ",ttString[ca.Transmission])
	fmt.Println("fuel:",ftStringEN[ca.FuelType])
	fmt.Println("year: ",ca.Year)
	fmt.Println("mileage:",ca.Mileage)
	fmt.Println("price: ", ca.Price)
	fmt.Println("found via: ",ca.Source)
	fmt.Println("link: ",ca.Link)
}