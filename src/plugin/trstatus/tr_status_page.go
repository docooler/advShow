package main

import (
	"TR"
	"fmt"
)

func InitData() []TR.TrPerTeam {
	teams := TR.Init()
	var tms = []TR.TrPerTeam{}
	for _, tm := range teams {
		tm.InitTRInfo()
		//do not display the zero  tr team
		if tm.TtlTrNr != 0 {
			tms = append(tms, tm)
		}
	}
	return tms
}
func createPage() {

}

func main() {
	tms := InitData()
	for _, tm := range tms {
		fmt.Print(tm)
	}
	fmt.Println("end")
}
