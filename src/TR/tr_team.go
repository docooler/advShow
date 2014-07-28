package TR

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "os"
	"regexp"
	"strings"
)

var teams = [...]string{"Rainbow", "Rainbow", "Rainbow", "Rainbow", "Rainbow"}
var mh = "radio_product_mhos"

type TrPerTeam struct {
	TeamName string
	ExTrNr   int16
	MtiTrNr  int16
	IntTrNr  int16
	TtlTrNr  int
	Trs      []*TrInfo
}

func Init() []TrPerTeam {
	size := len(teams)
	tm := make([]TrPerTeam, size)
	for i := 0; i < size; i++ {
		tm[i].TeamName = teams[i]
	}
	return tm
}

func (t *TrPerTeam) InitTRInfo() error {
	urls := get_tr_urls(t.TeamName)
	t.TtlTrNr = len(urls)
	t.Trs = make([]*TrInfo, t.TtlTrNr)
	for i, url := range urls {
		trinfo := NewTrInfo(url)
		trinfo.Init()
		t.count_tr_status(trinfo.Level)
		t.Trs[i] = trinfo
	}
	return nil
}
func (t *TrPerTeam) count_tr_status(trlevel int16) {
	switch trlevel {
	case External_TR:
		t.ExTrNr += 1
	case MTI_TR:
		t.MtiTrNr += 1
	case Internal_TR:
		t.IntTrNr += 1
	}
}

func get_team_status_html(mh string, teamName string) (string, error) {

	v := url.Values{}

	v.Set("bucket", mh)

	v.Add("team", teamName)

	v.Add("Search", "Show")

	request, err := http.NewRequest("POST", "http://lmr-radiosw-tr-tool.rnd.ki.sw.ericsson.se/tr_progress/per-team.php", strings.NewReader(v.Encode()))

	if err != nil {

		log.Println("Fatal error get_tearm_status", err)

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

		log.Println("get_tearm_status client Error ", err)

	}

	if response.StatusCode == 200 {

		body, _ := ioutil.ReadAll(response.Body)

		return string(body), nil

	}

	return "", nil

}

func get_team_status_html_stub(mh string, teamName string) (string, error) {

	html, err := ioutil.ReadFile("Rainbow.html")

	return string(html), err

}

func get_tr_per_team(teamName string, html string) string {
	splitStr := "<TR><TD bgcolor=#a0a0a0>" + teamName + "</TD>"
	splitTail := "</TABLE></body></html>"

	tables := strings.Split(html, splitStr)
	trs := strings.Split(tables[1], splitTail)
	tr := trs[0]
	return splitStr + tr
}

func get_tr_urls(teamName string) (urls []string) {

	html, err := get_team_status_html_stub("radio_product_mhos", teamName)
	if err != nil {
		CheckAndExit(err, "get_team_status_html_stub")
	}
	table := get_tr_per_team(teamName, html)
	targets := strings.Split(table, " target='_blank'>")
	re := regexp.MustCompile("https:.*\"")
	for _, target := range targets {
		if re.MatchString(target) {
			match := re.FindStringSubmatch(target)
			url := match[0]
			url = strings.Split(url, "\"")[0]
			urls = append(urls, url)
			fmt.Println(url)
		}
	}
	return
}
