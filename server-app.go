package main;

import (
	"net/http"
	"encoding/json"
	"fmt"
);

type Movie struct {
	Name string `json:"name"`
	Director string `json:"director"`
	Netflix bool `json:"netflix"`
}

func main() {
	localport := ":3000";
	fmt.Println("Server run in port 3000");

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "index.html");	
	});

	http.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		saga := Movie{"Black Mirror", "Charlie Brooker", true}
		json.NewEncoder(res).Encode(saga);
	})

	http.ListenAndServe(localport, nil);
}