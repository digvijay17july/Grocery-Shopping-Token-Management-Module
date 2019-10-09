package api

import (
	"Grocery-Shopping-Token-Management-Module/app/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()

}

func (a *App) setRouters(){
	fmt.Println("initializing request")

	a.Get("/download/{id}/{type}", a.getFile)
	a.Get("/getToken", a.GetToken)
	a.Get("/", a.init)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) GetToken(w http.ResponseWriter, r *http.Request) {
	service.GenerateFunction(w,r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App)init(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<html><h1>This is a SAMPLE HTML Page For EPFO</h1><br><h2>Hosted on Openshift</h2></html>")
}

func (a *App)getFile(w http.ResponseWriter, r *http.Request){
	service.DownloadFunction(w,r)
}
