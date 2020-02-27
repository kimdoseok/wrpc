// wrpc project main.go
package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ReturnValue struct {
	Errno  int
	Errmsg string
	Values interface{}
}

func getPassKeys(userid string) []string {
	outslice := make([]string, 3)
	h := md5.New()
	pstr := "Love means not ever having to say you are sorry."
	t := time.Now()
	t.AddDate(0, 0, -1)
	for i := -1; i < 2; i++ {
		tstr := fmt.Sprintf("%s%02d/%02d/%04d%s", userid, t.Month(), t.Day(), t.Year(), pstr)
		io.WriteString(h, tstr)
		outslice = append(outslice, fmt.Sprintf("%x", h.Sum(nil)))
		t.AddDate(0, 0, 1)
	}
	fmt.Println(outslice)
	return outslice
}

func runCommand(userid string, cmds []string, keys []string, w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "%v\n", cmds)
	if len(cmds) > 1 {
		if cmds[0] == "doc" {
			mod := Docs{userid, cmds, keys, w, r}
			mod.Run()
		} else if cmds[0] == "rpc" {
			mod := Rpcs{userid, cmds, keys, w, r}
			mod.Run()
		} else if cmds[0] == "pos" {
			mod := Poss{userid, cmds, keys, w, r}
			mod.Run()
		} else {

		}

	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path[1:], "/")
	if len(paths) < 2 {
		var outdata ReturnValue
		outdata.Errno = 1
		outdata.Errmsg = "Error! not sufficient parameters!\nFormat will be like userid/cmds/keys\n"
		outdata.Values = ""
		outbytes, _ := json.Marshal(outdata)
		fmt.Fprintf(w, string(outbytes))
		return
	}
	fmt.Println(getPassKeys(paths[0]))
	if r.Method != "POST" {
		var outdata ReturnValue
		outdata.Errno = 2
		outdata.Errmsg = "Error! It takes only post variables\n"
		outdata.Values = ""
		outbytes, _ := json.Marshal(outdata)
		fmt.Fprintf(w, string(outbytes))
		return
	}
	err := r.ParseForm()
	if err != nil {
		var outdata ReturnValue
		outdata.Errno = 3
		outdata.Errmsg = "Error! There is failed parseing form\n"
		outdata.Values = ""
		outbytes, _ := json.Marshal(outdata)
		fmt.Fprintf(w, string(outbytes))
		return
	}
	fmt.Println("r.PostForm:", r.PostForm)
	//fmt.Println("r.Form:", r.Form)
	userid := strings.ToUpper(paths[0])
	passkey := r.FormValue("p")
	passkeys := getPassKeys(paths[0])
	//fmt.p
	found := false
	for _, k := range passkeys {
		fmt.Println(passkey, k)
		if passkey == k {
			found = true
			break
		}
	}

	if !found {
		var outdata ReturnValue
		outdata.Errno = 4
		outdata.Errmsg = fmt.Sprintf("Error! passkey doesn't match!\nYour passkey is %s\n", passkey)
		outdata.Values = ""
		outbytes, _ := json.Marshal(outdata)
		fmt.Fprintf(w, string(outbytes))
		return
	}

	//r.ParseMultipartForm(32 << 20)
	//if err := r.ParseForm(); err != nil {
	//	fmt.Fprintf(w, "ParseForm() err: %v", err)
	//	return
	//}
	//fmt.Fprintf(w, "%v\n", r.Form)
	//for k, v := range r.Form {
	//	fmt.Fprintf(w, "key: %v, val: %v\n", k, strings.Join(v, ""))
	//}
	//fmt.Fprintf(w, "%v %v\n", reflect.TypeOf(r.Body), r.Body)
	cmds := strings.Split(paths[1], ".")
	keys := paths[2:]
	//fmt.Println(userid, cmds, keys, w, r)
	runCommand(userid, cmds, keys, w, r)

	//dec := json.NewDecoder(r.Body)
	//var v map[string]interface{}
	//if err := dec.Decode(&v); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//r.ParseForm()

	//for k := range v {
	//	fmt.Fprintf(w, "%s: %s\n", k, v[k])
	//}
	//if err := enc.Encode(&v); err != nil {
	//fmt.Println(err)
	//}

	//body, err := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%v\n", r.Form)
	//fmt.Fprintf(w, "Hi There, it's time to run module of %s\n", paths[2])
	//fmt.Fprintf(w, "%v %v %v", getPassKeys(paths[0]), len(paths[1]), len(getPassKeys(paths[0])))

	//file, handler, err := r.FormFile("f")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Fprintf(w, "file: %v\nhandle: %v\n", file, handler)
	//fmt.Fprintf(w, "filename: %v, header: %v, size: %v\n", handler.Filename, handler.Header, handler.Size)
	//for k, v := range handler.Header {
	//	fmt.Fprintf(w, "key: %v, val: %v\n", k, v)
	//	for k1, v1 := range v {
	//		fmt.Fprintf(w, "key1: %v, val1: %v\n", k1, v1)
	//}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
