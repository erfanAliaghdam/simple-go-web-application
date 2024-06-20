package main

import (
	"fmt"
	"net/http"
)


func (app *application) Index(
	w http.ResponseWriter, 
	r *http.Request){
	fmt.Fprint(w, "Hey, how u doin?!")
}