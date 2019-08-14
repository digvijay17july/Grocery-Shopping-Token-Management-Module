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
	fmt.Println("Port No. is :"+port)
	app.Run(":"+port)
}