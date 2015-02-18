package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
)

func doGrep(regex string) {
	url := "http://127.0.0.1:9999/" + regex
	fmt.Println(url)
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}

func main() {
	app := cli.NewApp()
	app.Name = "cagrep_client"
	app.Usage = ""
	app.Version = "0.0.1"
	app.Author = "sago35"
	app.Email = "sago35@gmail.com"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) > 0 {
			doGrep(c.Args()[0])
		}
	}
	app.Run(os.Args)
}
