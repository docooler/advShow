package main

import (
	"TR"
	// "fmt"
	"adutils"
	"html/template"
	"os"
)

const (
	STATIC_DIR  = "../static/"
	VIEW_DIR    = "../view/"
	TMP_FILE    = STATIC_DIR + "tmp.html"
	TARGET_FILE = STATIC_DIR + "tr_chart.html"
)

type PageInfo struct {
	MaxTRNR int
	TrTeams []TR.TrPerTeam
}

func InitData() []TR.TrPerTeam {
	teams := TR.Init()
	var tms = []TR.TrPerTeam{}
	for _, tm := range teams {
		tm.InitTRInfo(0)
		//do not display the zero  tr team
		if tm.TtlTrNr != 0 {
			tms = append(tms, tm)
		}
	}
	return tms
}
func moveToTarget() {

	flag, err := adutils.Exists(TARGET_FILE)
	if flag {
		err = adutils.Unlink(TARGET_FILE)
		if err != nil {
			os.Exit(1)
			return
		}
	}
	adutils.Rename(TMP_FILE, TARGET_FILE)
}
func createPage() {
	var page PageInfo
	flag, err := adutils.Exists(TMP_FILE)
	if flag {
		err = adutils.Unlink(TMP_FILE)
		if err != nil {
			os.Exit(1)
			return
		}
	}

	filefd, err := os.Create(TMP_FILE)
	defer filefd.Close()
	if err != nil {
		os.Exit(1)
	}

	tms := InitData()

	page.TrTeams = tms
	max_tr_num := 0
	for _, tm := range tms {
		// fmt.Print(tm)
		if max_tr_num < tm.TtlTrNr {
			max_tr_num = tm.TtlTrNr
		}
	}
	max_tr_num *= 120
	max_tr_num /= 100
	page.MaxTRNR = max_tr_num
	t, _ := template.ParseFiles(VIEW_DIR + "tr_chart.html")
	t.Execute(filefd, page)
}

func main() {
	createPage()
	moveToTarget()
}

