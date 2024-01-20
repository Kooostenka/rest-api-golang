package main

import (
	"app/internal/user"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("hello %s", name)))
}

func main() {
	fmt.Println("create router")
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {

	listner, err := net.Listen("tcp", "127.0.0.1:1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listner))
}
