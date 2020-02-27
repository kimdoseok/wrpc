package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Docs struct {
	userid string
	cmds   []string
	keys   []string
	w      http.ResponseWriter
	r      *http.Request
}

func (d Docs) Run() {
	if len(d.cmds) < 2 {
		return
	} else {
		cmd := strings.ToLower(d.cmds[1])
		fmt.Print(d.cmds)
		if strings.Compare(cmd, "a") > 0 {
			d.Read()
		}
	}
}

func (d Docs) Read() {
	fmt.Fprintln(d.w, "Read")
}
