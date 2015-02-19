package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
)

func queryGrep(c *cli.Context) {
	regex := c.Args()[0]
	url := "http://127.0.0.1:" + c.String("port") + "/" + regex

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
