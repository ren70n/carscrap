package main

import (
	"log"
	"github.com/ren70n/carscrap/adapters"
)

// I have to get the list of web services with used cars
// parse data trought adapters
// in search of the optimal car for me

// mainly to check what new has appeared (and what's disappeared)

func main(){
	// first version is just to scrap autohero.pl
	log.Println("main is running")
	adapters.AutoheroGetter()
}

