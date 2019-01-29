package main

import (
	"log"
	"net/http"

	"./controllers"
)

func main() {
	port := ":8080"
	//controllers.Login()
	//controllers.run()
	http.HandleFunc("/", controllers.Login)
	err := http.ListenAndServe(port, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
// 	// attention: If you do not call ParseForm method, the following data can not be obtained form
// 	fmt.Println(r.Form) // print information on server side.
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello world!") // write data to response
// }
