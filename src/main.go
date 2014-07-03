package main 

import (
         "net/http"
         "log" 
         "html/template"
         "fmt"
       )

func rootHandler(w http.ResponseWriter, r * http.Request) {
    r.ParseForm()
if r.Method == "GET" {
      t, err := template.ParseFiles("static/index.html")
      if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
      }

      t.Execute(w, nil )

      
      fmt.Println ("End root")
      return
   }
}

func displayHandler(w http.ResponseWriter, r * http.Request) {
    r.ParseForm()
    t, err := template.ParseFiles("FT_88_4_4FT_88_4_5.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    index := r.Form["index"]
    fmt.Println(index)
    t.Execute(w, nil )
    fmt.Println("End displayHandler")
    return
}


func main() {
   http.HandleFunc("/", rootHandler)
   http.HandleFunc("/display" , displayHandler)
  
   err := http.ListenAndServe(":8080", nil)
   if err != nil {
      log.Fatal("God Like listen wrong: ", err.Error())
   }
   
}
