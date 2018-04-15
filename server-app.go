package main;

import (
	"net/http"
	"fmt"
);

func main() {
	fmt.Println("Server run in port 3000");
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "index.html");	
	})

	http.ListenAndServe(":3000", nil);
}