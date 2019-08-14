package main

import api "Grocery-Shopping-Token-Management-Module/app/app"

func main(){
	app:= api.App{}
	app.Initialize()
	app.Run(":3011")
}