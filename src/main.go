package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

const ADV_LEN = 5

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("../static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)

		fmt.Println("End root")
		return
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getAdvFilename(index string) (string, error) {
	i, err := strconv.Atoi(index)
	if err != nil {
		// handle error
		fmt.Println(err)
		return " ", err
	}
	fmt.Println(index)

	findex := i % ADV_LEN
	filename := "../static/" + strconv.Itoa(findex) + ".html"

	ret, _ := exists(filename)
	if ret != true {
		fmt.Println(filename + "not exist in the file")
	}
	return filename, nil

}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index := r.Form["index"]

	filename, err := getAdvFilename(index[0])
	if err != nil {
		fmt.Println("displayHandler getAdvFilename failed")
		filename = "0.html"
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
	http.HandleFunc("/display", displayHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("God Like listen wrong: ", err.Error())
	}

}
