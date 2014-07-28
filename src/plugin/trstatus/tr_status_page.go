package main

import (
	"TR"
	"fmt"
)

func InitData() {
	teams := TR.Init()
	for i, tm := range teams {
		//do not display the zero  tr team
		if tm.TtlTrNr == 0 {
			delete teams[i]
		}


	}
}
func createPage() {

}

func main() {
	InitData()
	fmt.Println('end')
}
