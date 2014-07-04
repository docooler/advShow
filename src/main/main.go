package main

import (
	"adutils"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	STATIC_FILE = "../../static/"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	requestfile := STATIC_FILE + r.URL.Path

	ret, _ := adutils.Exists(requestfile)
	if ret {
		content, err := ioutil.ReadFile(requestfile)
		if err == nil {
			w.Write(content)
			return
		}
	}

	fmt.Println("request file do not exist . End root request .Do nothing")
	return
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("../../static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		srv, err := adutils.ServerParse()

		if err != nil {
			fmt.Println("Parse server file error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		srv.DspCtrl.Delaytime *= 1000
		fmt.Println(srv)

		t.Execute(w, srv)

		fmt.Println("End showHandler")
		return
	}
}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index := r.Form["index"]

	filename, linktype, err := adutils.GetAdvFilename(index[0])
	if err != nil {
		fmt.Println("displayHandler getAdvFilename failed")
		filename = "0.html"
	}
	if linktype == 1 {
		http.Redirect(w, r, filename, 300)
		return
		// resp, err := http.Get(filename)
		//   	defer resp.Body.Close()
		//   	if err != nil { panic(err) }
		//   	for k, v := range resp.Header {
		//   	    for _, vv := range v {
		//   	        w.Header().Add(k, vv)
		//   	    }
		//   	}

		//   	w.WriteHeader(resp.StatusCode)
		//   	result, err := ioutil.ReadAll(resp.Body)
		//   	if err != nil  { panic(err) }
		//   	w.Write(result)
		//   	return
	}

	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

	fmt.Println("End displayHandler")
	return
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/show", showHandler)
	http.HandleFunc("/display", displayHandler)

	srv, err := adutils.ServerParse()

	if err != nil {
		log.Fatal("server.xml error", err.Error())
		return
	}
	port := ":" + srv.Monitor.Port
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("God Like listen wrong: ", err.Error())
	}

}
