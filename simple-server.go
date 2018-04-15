package main;

import "net/http";
import "io";
import "fmt";

func main() {

	fmt.Println("Server Runing!");
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil);	
}

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Println("New petition");
	io.WriteString(res, "<h1>Welcome a your server golang</h1>");
}
