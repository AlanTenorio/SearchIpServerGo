package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a commandLIne
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CommandLIneGo"
	app.Usage = "Seach Ip and Name Server"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "...",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP",
			Flags:  flags,
			Action: seachIps,
		},
		{
			Name:   "servers",
			Usage:  "Serch IPs on Internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func seachIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		fmt.Printf("Error: Could not find the IP '%s'\n", host)
		return
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		fmt.Printf("Error: Could not find the host '%s'\n", host)
	}

	for _, server := range servers {
		fmt.Printf(server.Host)
		return
	}

}
