package main

import (
       
       "log"
    )

// main is the entry point of the application. It initializes the API server and starts it on port 3000.
func main() {

	store,err:=NewPostgresStore()
	if err!=nil{
		log.Fatal(err)
	}
       if err:=store.Init();err!=nil{
		log.Fatal(err)
	   }
	// Create a new API server instance listening on port 3000.
	server := NewApiServer(":3000",store )

	 //Run the server to start listening for requests.
	server.Run()
}