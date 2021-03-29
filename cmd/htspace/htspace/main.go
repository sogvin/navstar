// Command htspace runs the htspace.Application as a web server.
package main

import (
	"log"
	"net/http"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/spaceflight"
	"github.com/gregoryv/spaceflight/cmd/htspace"
	"github.com/gregoryv/wolf"
)

func main() {
	run(wolf.NewOSCmd())
}

func run(cmd wolf.Command) {
	var (
		cli  = cmdline.NewParser(cmd.Args()...)
		help = cli.Flag("-h, --help")
		bind = cli.Option("-b, --bind").String(":9188")
	)

	log.SetFlags(0)

	switch {
	case help:
		cli.WriteUsageTo(cmd.Stdout())

	case !cli.Ok():
		log.Println(cli.Error())
		cmd.Exit(1)

	default:
		sys := spaceflight.NewSystem()
		app := htspace.NewApplication(sys)
		log.Println("listening on", bind)
		err := http.ListenAndServe(bind, app.Router())
		if err != nil {
			log.Println(err)
			cmd.Exit(1)
		}
	}
}
