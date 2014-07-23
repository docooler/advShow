package main

import (
        "log"
        "net/http"
	    "net/url"
	    "strings"
        "io/ioutil"
        )

func get_tearm_status(mh string, teamName string) (string, error){
    v := url.Values{}
    v.Set("bucket", mh)
    v.Add("team", teamName)
    v.Add("Search", "Show")
    request, err := http.NewRequest("POST", "http://lmr-radiosw-tr-tool.rnd.ki.sw.ericsson.se/tr_progress/per-team.php", strings.NewReader(v.Encode()))
    if err != nil {
        log.Println("Fatal error get_tearm_status", err.Error)
        return "", err
    }
    request.Header.Set("Host", "lmr-radiosw-tr-tool.rnd.ki.sw.ericsson.se")
    request.Header.Set("Connection", "keep-alive")
    request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    request.Header.Set("Origin", "http://lmr-radiosw-tr-tool.rnd.ki.sw.ericsson.se")
    request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36" )
    request.Header.Set("Referer", "http://lmr-radiosw-tr-tool.rnd.ki.sw.ericsson.se/tr_progress/per-team.php")
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    client := &http.Client{}
    response, err := client.Do(request)
    defer response.Body.Close()

    if err != nil {
        log.Println("get_tearm_status client Error ", err.Error)    
    }

    if response.StatusCode == 200 {
        body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
    }

    return "", nil
}
func main() {

    body, err := get_tearm_status("radio_product_mhos", "Rainbow")
    if err != nil {
        log.Fatal("main get_tearm_status error ", err.Error)
        return
    }
    log.Println(body)
}
