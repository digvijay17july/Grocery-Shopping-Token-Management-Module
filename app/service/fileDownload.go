package service

import (
	"github.com/gorilla/mux"
	"net/http"
)

func DownloadFunction(w http.ResponseWriter,r * http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]
	typeOfFile := vars["type"]

	//fs := http.FileServer(http.Dir("../static/"+id))

   if typeOfFile=="png" {
   	w.Header().Set("Content-Type","image/png")
   } else {
		w.Header().Set("Content-Type","text/css")
	}
	w.WriteHeader(200)
	//http.Handle("/files/", http.StripPrefix("/files", fs))
	http.ServeFile(w,r,"static/"+id+"."+typeOfFile)

}