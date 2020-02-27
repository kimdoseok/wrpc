// rpcs
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Poss struct {
	userid string
	cmds   []string
	keys   []string
	w      http.ResponseWriter
	r      *http.Request
}

func (d Poss) Run() {
	if len(d.cmds) < 2 {
		return
	} else {
		cmd := strings.ToLower(d.cmds[1])
		fmt.Println("cmd:" + cmd)
		funcs := map[string]func(){
			"a": d.GetItems,
			"b": d.GetItem,
			"c": d.PutItem,
			"d": d.Test}
		found := false
		for k := range funcs {
			if cmd == k {
				funcs[k]()
				found = true
				break
			}
		}
		if !found {
			var outdata ReturnValue
			outdata.Errno = 10
			outdata.Errmsg = fmt.Sprintf("Error! empty results\n")
			outdata.Values = ""
			outbytes, _ := json.Marshal(outdata)
			fmt.Fprintf(d.w, string(outbytes))
		}
	}
}

func (d Poss) GetItems() {
	var iit Iit
	InitDB()
	keystr := ""
	for _, k := range d.keys {
		keystr = keystr + " " + k
	}
	keystr = strings.TrimSpace(keystr)
	iits := iit.GetItems(keystr)
	//fmt.Fprintln(d.w, iits)
	jsoniit, err := json.Marshal(iits)
	if err != nil {
		fmt.Println("Json Mashall Error!")
	}
	fmt.Fprintln(d.w, string(jsoniit))
	//fmt.Println(len(jsoniit))
}

func (d Poss) GetItem() {
	var iit Iit
	InitDB()
	keystr := ""
	for _, k := range d.keys {
		keystr = keystr + " " + k
	}
	keystr = strings.TrimSpace(keystr)
	iits := iit.GetItem(keystr)
	//fmt.Fprintln(d.w, iits)
	jsoniit, err := json.Marshal(iits)
	if err != nil {
		fmt.Println("Json Mashall Error!")
	}
	fmt.Fprintln(d.w, string(jsoniit))
	//fmt.Println(len(jsoniit))
}

func (d Poss) PutItem() {
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

func (d Poss) Test() {
	//InitDB()
	//iits := GetAllItems()
	//fmt.Fprintln(d.w, iits)
	//jsoniit, err := json.Marshal(iits[0])
	//if err != nil {
	//	fmt.Println("Json Mashall Error!")
	//}
	//fmt.Fprintln(d.w, jsoniit)
	for k, v := range d.r.Form {
		//fmt.Printf(d.w, "key: %v, val: %v\n", k, strings.Join(v, ""))
		fmt.Printf("key: %v, val: %v\n", k, strings.Join(v, ""))
	}
	var jdata map[string]interface{}
	jbytes := []byte(d.r.Form["j"][0])
	err := json.Unmarshal(jbytes, &jdata)
	if err != nil {
		fmt.Fprintf(d.w, "Error!")
	}
	var outdata ReturnValue
	outdata.Errno = 0
	outdata.Errmsg = "OK"
	jdata["p"] = d.r.Form["p"][0]
	outdata.Values = jdata
	for k, v := range jdata {
		fmt.Printf("Jdata key: %v, val: %v\n", k, v)
	}
	//fmt.Println(jdata)
	outdatabytes, _ := json.Marshal(outdata)
	fmt.Fprintf(d.w, string(outdatabytes))
	fmt.Printf(string(outdatabytes))

}
