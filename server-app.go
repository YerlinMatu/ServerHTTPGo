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
	localport := ":8080";
	fmt.Printf("Server run in port %s", localport);

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "index.html");	
	});

	http.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		saga := Movie{"Black Mirror", "Charlie Brooker", true}
		json.NewEncoder(res).Encode(saga);
	})

	http.ListenAndServe(localport, nil);
}