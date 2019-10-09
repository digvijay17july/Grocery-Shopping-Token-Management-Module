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
	fmt.Fprint(w,"<html lang='en'><head>
	<meta charset='utf-8'>
	<title>Deployment Demonstration</title>
	<meta name='viewport' content='width=device-width, initial-scale=1.0'>
	<style>
	  HTML{height:100%;}
	  BODY{font-family:Helvetica,Arial;display:flex;display:-webkit-flex;align-items:center;justify-content:center;-webkit-align-items:center;-webkit-box-align:center;-webkit-justify-content:center;height:100%;}
	  .box{background:#b5d4a8;color:white;text-align:center;border-radius:10px;display:inline-block;}
	  H1{font-size:10em;line-height:1.5em;margin:0 0.5em;}
	  H2{margin-top:0;}
	</style>
  </head>
  <body>
  <div class='box'><h1>This is a SAMPLE HTML Page For EPFO</h1><h2>Hosted on Openshift</h2></div>
  
  </body></html>")
}

func (a *App)getFile(w http.ResponseWriter, r *http.Request){
	service.DownloadFunction(w,r)
}
