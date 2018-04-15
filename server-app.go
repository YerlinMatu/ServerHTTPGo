package main;

import (
	"net/http"
	"fmt"
);

func main() {
	localport := ":3000";
	fmt.Println("Server run in port 3000");
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "index.html");	
	});

	http.ListenAndServe(localport, nil);
}