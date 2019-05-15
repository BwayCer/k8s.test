package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	strTimeMs := time.Now().Format(".000000")[1:]
	intTimeMs, _ := strconv.Atoi(strTimeMs)
	id = strconv.FormatUint(uint64(intTimeMs), 16)

	insCmdFlag := handleFlag()

	server := &http.Server{Addr: insCmdFlag.Addr}

	http.HandleFunc("/favicon.ico", func(resp http.ResponseWriter, requ *http.Request) {
		resp.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/panic", func(resp http.ResponseWriter, requ *http.Request) {
		lxoxg("panic")
		resp.WriteHeader(http.StatusInternalServerError)
		panic(server.Shutdown(context.TODO()))
	})

	http.HandleFunc("/", func(resp http.ResponseWriter, requ *http.Request) {
		urlPath := requ.URL.Path
		lxoxg("cheer " + urlPath)
		pageHtml := strings.Replace(cheerPageTmpl, "{{.path}}", urlPath, 1)
		resp.Header().Set("Content-Type", "text/html")
		resp.Header().Set("Content-Length", strconv.Itoa(len(pageHtml)))
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, pageHtml)
	})

	panic(server.ListenAndServe())
}

type CmdFlag struct {
	Addr string
}

var id string

var cheerPageTmpl = `<html>
    <head>
        <meta charset="utf-8">
        <title>Cheer</title>
    </head>
    <body>
        <h2>Cheer {{.path}}</h2>
        <div>go to <a href="/panic">/panic</a>, and stop this machine.</div>
    </body>
</html>`

func handleFlag() CmdFlag {
	insCmdFlag := CmdFlag{}
	flag.StringVar(&insCmdFlag.Addr, "addr", "0.0.0.0:8080", "http service address.")
	flag.Parse()
	return insCmdFlag
}

func lxoxg(txt string) {
	now := time.Now()
	fmt.Println(now.Format("15:04:05.000000"), "(id: "+id+")", txt)
}
