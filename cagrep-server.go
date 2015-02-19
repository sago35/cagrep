package main

import (
	"bufio"
	"code.google.com/p/mahonia"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/hoisie/web"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var dict = []string{
	"hello",
	"world",
}

type File struct {
	Filename string
	Lines    []string
}

var files = []File{}

func doGrep(val string) []string {
	//i, _ := strconv.ParseInt(val, 0, 32)
	match := []string{}
	regex := regexp.MustCompile(val)

	for _, f := range files {
		for lineNum, line := range f.Lines {
			if regex.MatchString(line) {
				match = append(match, fmt.Sprintf("%s:%d:%s", f.Filename, lineNum+1, line))
			}
		}
	}

	return match
}

func appendFiles(ch chan File) {
	for c := range ch {
		fmt.Println(c.Filename)
		files = append(files, c)
	}
}

func hello(ctx *web.Context, val string) string {
	ctx.SetHeader("Content-Type", "text/plain", true)
	return strings.Join(doGrep(val), "\r\n")
}

func cagrepServer(c *cli.Context) {
	ch := make(chan File)

	root := `.`
	if len(c.Args()) > 0 {
		root = c.Args()[0]
	}
	go func() {
		err := filepath.Walk(root,
			func(path string, info os.FileInfo, err error) error {
				rel, _ := filepath.Rel(root, path)
				if info.IsDir() {
					return nil
				} else if regexp.MustCompile(`.git`).MatchString(path) {
					return nil
				} else if regexp.MustCompile(`\.(db|lib|rel|a|exe|xls)$`).MatchString(path) {
					return nil
				} else if regexp.MustCompile(`cscope\.out$`).MatchString(path) {
					return nil
				}

				// ここでファイルを読む

				f, err := os.Open(path)
				if err != nil {
					fmt.Fprintf(os.Stderr, "File open error (%s) : %v", path, err)
					os.Exit(1)
				}

				scanner := bufio.NewScanner(f)
				lines := []string{}
				for scanner.Scan() {
					lines = append(lines, mahonia.NewDecoder("cp932").ConvertString(scanner.Text()))
				}

				ch <- File{
					Filename: rel,
					Lines:    lines,
				}
				return nil
			})
		if err != nil {
			panic(err)
		}
		close(ch)
	}()

	appendFiles(ch)

	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:" + c.String("port"))
}
