package main

import (
	"goweb/web22-1/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {

	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Print("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
