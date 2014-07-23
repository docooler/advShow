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
func get_team_status_stub(mh string, teamName string) (string, error) {
	html, err := ioutil.ReadFile("Rainbow.html")

	return string(html), err
}

func get_tr_status_page_head(mh string, teamName string) (string, error) {
	html, err := get_team_status_stub(mh, teamName)
	if err != nil {
		return "", err
	}
	splitStr := "<TR><TD bgcolor=#a0a0a0>" + teamName + "</TD>"
	splits := strings.Split(html, splitStr)
	return splits[0], nil
}
func get_tr_per_team(teamName string, html string) string {
	splitStr := "<TR><TD bgcolor=#a0a0a0>" + teamName + "</TD>"
	splitTail := "</TABLE></TD></TR>"

	tables := strings.Split(html, splitStr)
	trs := strings.Split(tables[1], splitTail)
	tr := trs[0]
	return splitStr + tr + splitTail
}

var teams = [...]string{"Rainbow", "Rainbow", "Rainbow", "Rainbow", "Rainbow"}

func creat_tr_status_page(filename string) error {
	mh := "radio_product_mhos"
	page, err := get_tr_status_page_head(mh, teams[0])
	if err != nil {
		ErrorAndExit(err)

	}
	for _, team := range teams {
		// log.Println(team)
		html, err := get_team_status_stub(mh, team)
		if err != nil {
			ErrorAndExit(err)
		}
		tr := get_tr_per_team(team, html)
		page = page + tr
	}
	page = page + "</TABLE></body></html>"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(page)
	if err != nil {
		panic(err)
	}

	return nil

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
	// html, err := ioutil.ReadFile("Rainbow.html")

	// if err != nil {
	// 	ErrorAndExit(err)
	// }

	// str := get_tr_per_team("Rainbow", string(html))
	// log.Println(str)
	// return
	creat_tr_status_page("tmp.html")

}
