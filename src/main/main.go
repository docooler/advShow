package main

import (
	"adutils"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	STATIC_DIR = "../static/"
    VIEW_DIR   = "../view/"
)
type home struct{
    Title string
} 
func ext2Mime(ext string) string {
    switch ext {
    case ".css":
        return "text/css"
    case ".js":
        return "text/js"
    case ".html":
        return "text/html"
    default:
        return ""
    }
    return "*/*"

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
    log.Println("Info:request file ", r.URL.Path)

    //visitor main page
    if r.URL.Path == "/" {
        title := home{Title: "advise show system "}
        t, _ := template.ParseFiles(VIEW_DIR + "index.html")
        t.Execute(w, title)
        return 
    }

    requestfile := STATIC_DIR + r.URL.Path
    
    ret,_ := adutils.Exists(requestfile)
    if ret {
        content, err := ioutil.ReadFile(requestfile)
        if err == nil {
            mType := ext2Mime(requestfile)
            if mType == "text/js" {
                w.Header().Set("Content-Type", mType)
            }
            w.Header().Set("Cache-Control", "public, max-age=86400")

            w.Write(content)
            return 
        }
    }
    
    http.NotFound(w, r)
    log.Println("Error: file not find. path=" , r.URL.Path)
    return

}

func showHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles(VIEW_DIR + "showCtrl.html")
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

func fileHandler(w http.ResponseWriter, r *http.Request) {
        http.StripPrefix("/file", http.FileServer(http.Dir(STATIC_DIR))).ServeHTTP(w, r)    
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
    http.HandleFunc("/file", fileHandler)

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
