// rpcs
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Rpcs struct {
	userid string
	cmds   []string
	keys   []string
	w      http.ResponseWriter
	r      *http.Request
}

func (d Rpcs) Run() {

	if len(d.cmds) < 2 {
		return
	} else {
		cmd := strings.ToLower(d.cmds[1])
		if cmd == "a" {
			d.Get()
		} else if cmd == "b" {
			d.Put()
		}

	}
}

func (d Rpcs) Get() {
	a := d.r.FormValue("a")
	fmt.Fprintf(d.w, "Run first method in rpc struc. a=%v\n", a)
	fmt.Fprintln(d.w, "Get")
}

func (d Rpcs) Put() {
	fmt.Fprintln(d.w, "Put")
	d.r.ParseMultipartForm(32 << 20)
	if err := d.r.ParseForm(); err != nil {
		fmt.Fprintf(d.w, "ParseForm() err: %v", err)
		return
	}

	file, handler, err := d.r.FormFile("f")
	fmt.Fprintf(d.w, "file: %v\nhandle: %v\n", file, handler)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
