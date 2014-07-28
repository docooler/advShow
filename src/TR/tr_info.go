package TR

import (
	// "fmt"
	// "io/ioutil"
	"log"
	"os"
	// "net/http"
	// "net/url"
	// "regexp"
	// "strings"
)

const (
	External_TR = iota
	MTI_TR
	Internal_TR
)

type TrInfo struct {
	Url     string
	Level   int16
	OwnerId string
	email   string
	expire  int16
}

func NewTrInfo(url string) *TrInfo {
	return &TrInfo{
		Url: url,
	}
}
func (tr *TrInfo) Init() error {
	html, err := getHtmlPage(tr.Url)
	CheckAndExit(err, "getHtmlPage "+tr.Url)
	tr.Level = parseLevel(html)
	tr.OwnerId = parseOwner(html)
	tr.email = getEmail(tr.OwnerId)
	return nil
}

//TODO
func parseLevel(html string) int16 {
	return MTI_TR
}

//TODO
func parseOwner(html string) string {
	return "elaiyan"
}

//TODO
func getEmail(ownerId string) string {
	return "laiyuan.yang@ericsson.com"
}

//TODO
func getHtmlPage(url string) (string, error) {
	return "", nil
}

//TODO
func CheckAndExit(err error, addInfo string) {
	if nil != err {
		log.Println(addInfo, err)
		os.Exit(1)
	}
}
