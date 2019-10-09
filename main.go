package main

import (
	api "Grocery-Shopping-Token-Management-Module/app/app"
	"fmt"
	"os"
)

func main(){
	app:= api.App{}
	app.Initialize()
	app.Run(":8080")
}
