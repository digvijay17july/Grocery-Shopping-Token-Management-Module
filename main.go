package main

import (
	api "Grocery-Shopping-Token-Management-Module/app/app"
	"fmt"
	"os"
)

func main(){
	app:= api.App{}
	app.Initialize()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	app.Run(":"+port)
	
	
}
