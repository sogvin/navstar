// Command starplan exposes the navstar/htapi through a TCP server
package main

import (
	"log"
	"net/http"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/navstar"
	"github.com/gregoryv/navstar/htapi"
)

func main() {
	var (
		cli  = cmdline.NewBasicParser()
		bind = cli.Option("-b, --bind").String(":9188")
	)
	cli.Parse()

	log.SetFlags(0)

	sys := navstar.NewSystem()
	router := htapi.NewRouter(sys)
	log.Println("listening on", bind)
	err := http.ListenAndServe(bind, router)
	if err != nil {
		log.Println(err)
	}
}
