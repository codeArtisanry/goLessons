//Package main is main package of web server and contains main function
package main

import (
	"log"
	"net/http"
	"github.com/tahasevim/webserver/profilehandler"
	"flag"
)

func main() {
	p := flag.String("port","8080","holds port") //port flag
	flag.Parse()
	myServer := http.Server{
		Addr:	":"+ *p,//address that imply localhost
		Handler: profilehandler.NewProfileHandler(),//handler of webserver
	}
	log.Println("Server started at port: "+ *p)
	log.Println(myServer.ListenAndServe())	
}
