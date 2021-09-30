package main

import "github.com/urfave/cli/v2"

type flagArgs struct {
	url         string
	name        string
	num         int
	boilerplate string
	path        string
}

func extractFlagArgs(c *cli.Context) *flagArgs {
	url := c.String(URL)
	num := c.Int(NUMBER)
	name := c.String(NAME)
	boilerplate := c.String(BOILERPLATE)
	path := c.String(LOCATION)

	return &flagArgs{
		url:         url,
		name:        name,
		num:         num,
		boilerplate: boilerplate,
		path:        path,
	}
}
