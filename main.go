package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

// lib-src [--source|--homepage|--bugtracker] (gem|npm|cpan) <name>

type libInfo struct {
	source     string
	homepage   string
	bugtracker string
}

type infoFetcher func(libName string) (*libInfo, error)

func siteAction(fetcher infoFetcher) func(c *cli.Context) {
	return func(c *cli.Context) {
		info, err := fetcher(c.Args().First())
		if err != nil {
			panic(err)
		}

		if c.GlobalBool("source") {
			fmt.Println(info.source)
		} else if c.GlobalBool("homepage") {
			fmt.Println(info.homepage)
		} else if c.GlobalBool("bugtracker") {
			fmt.Println(info.bugtracker)
		} else {
			fmt.Println(info.source)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "lib-src"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "source",
			Usage: "Show source",
		},
		cli.BoolFlag{
			Name:  "homepage",
			Usage: "Show homepage",
		},
		cli.BoolFlag{
			Name:  "bugtracker",
			Usage: "Show bug tracker",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "gem",
			Action: siteAction(fetchRubyGems),
		},
		cli.Command{
			Name:   "npm",
			Action: siteAction(fetchNPM),
		},
		cli.Command{
			Name:   "cpan",
			Action: siteAction(fetchCPAN),
		},
	}
	app.Run(os.Args)
}
