package main

import (
	"log"
	"net/http"

	handler "github.com/cache/src/handler"
	lib "github.com/cache/src/lib"
	util "github.com/cache/src/util"
)

func main() {
	log.Println("Start this project")

	go lib.ListenChannel(util.RegisterTask()...)

	http.HandleFunc("/read", handler.Read)
	http.HandleFunc("/write", handler.Write)
	http.ListenAndServe(":8080", nil)
}
