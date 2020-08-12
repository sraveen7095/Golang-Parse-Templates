package main

import(
"os"
"net/http"
"log"
"controller"
)

func main(){
port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()

	
    log.Println("Server started on: http://localhost:8080")
    mux.HandleFunc("/",controller.Index)
    
	http.ListenAndServe(":"+port, mux)
	
}