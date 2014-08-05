package main

import (
	"TR"
	"adutils"
	"log"
	"time"
)

func InitData() []TR.TrPerTeam {
	teams := TR.Init()
	var tms = []TR.TrPerTeam{}
	for _, tm := range teams {
		tm.InitTRInfo(1)
		//do not display the zero  tr team
		if tm.TtlTrNr != 0 {
			tms = append(tms, tm)
		}
	}
	return tms
}

func genMailList() map[string]*MailInfo {
	mailMap := make(map[string]*MailInfo)

	tms := InitData()
	for _, tm := range tms {
		// log.Println("TrNr per team", tm.TeamName, tm.TtlTrNr)
		// log.Println("tr length ", len(tm.Trs))
		for _, trInfo := range tm.Trs {
			// log.Println("genMailList trInfo: ", trInfo)
			if mailInfo, ret := mailMap[trInfo.Email]; ret {
				// log.Println("genMailList append to mail list")
				// log.Println("genMailList len:", len(mailInfo.Trs))
				mailInfo.Trs = append(mailInfo.Trs, trInfo.Clone())
			} else {
				mI := NewMailInfo(trInfo)
				mailMap[trInfo.Email] = mI
			}
		}
	}
	return mailMap
}
func composeMail(mailList map[string]*MailInfo) {
	return

}
func composeAndSendMail(mailList map[string]*MailInfo) {
	for _, v := range mailList {
		// log.Println(v.Email)
		// log.Println("total tr number :", len(v.Trs))
		// for _, trInfo := range v.Trs {
		// 	log.Println(trInfo)
		// }
		// log.Println(v.ComplseMail())
		v.ComposeMail()
		adutils.SendMail(v.Email, v.ContentFile)
		time.Sleep(time.Second * 10)
	}
}
func main() {
	mailList := genMailList()
	composeMail(mailList)
	composeAndSendMail(mailList)

	log.Println("Send mail done")
}
