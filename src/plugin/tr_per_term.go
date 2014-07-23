package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func get_team_status(mh string, teamName string) (string, error) {
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
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36")
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

func get_tr_per_team(teamName string) string {
	splitStr := "<TR><TD bgcolor=#a0a0a0>" + teamName + "</TD>"
	splitTail := "</TABLE></TD></TR>"
	html, err := ioutil.ReadFile("Rainbow.html")

	if err != nil {
		ErrorAndExit(err)
	}

	tables := strings.Split(string(html), splitStr)
	// log.Println(tables)
	trs := strings.Split(tables[1], splitTail)
	tr := trs[0]
	return splitStr + tr + splitTail
}

func ErrorAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {

	// body, err := get_team_status("radio_product_mhos", "Rainbow")
	// if err != nil {
	//     log.Fatal("main get_tearm_status error ", err.Error)
	//     return
	// }
	// log.Println(body)
	str := get_tr_per_team("Rainbow")
	log.Println(str)
	return

}
