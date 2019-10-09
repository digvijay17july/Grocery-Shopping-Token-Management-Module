package main

import (
	api "Grocery-Shopping-Token-Management-Module/app/app"
	"fmt"
	"os"
)

func main(){
	app:= api.App{}
	fmt.Print("Running 1")
	app.Initialize()

	fmt.Print("Running 2")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	fmt.Print("Running 3")
	app.Run(":"+port)
	
	fmt.Print("Running 4")
	
}
