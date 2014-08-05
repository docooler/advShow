package main

import (
	"TR"
	"adutils"
	"email"
	"log"
	"os"
)

const (
	CURRENT_DIR = "./"
)

type MailInfo struct {
	Email       string
	Name        string
	Subject     string
	ContentFile string
	Trs         []TR.TrInfo
}

func NewMailInfo(tr TR.TrInfo) *MailInfo {
	mi := &MailInfo{
		Email: tr.Email,
		Trs:   make([]TR.TrInfo, 1, 5),
	}
	mi.Trs[0] = tr.Clone()
	return mi
}

func (m *MailInfo) ComposeMail() string {
	e := email.NewEmail()
	e.From = "docooler@fep.loaddomain"
	e.To = []string{m.Email}
	e.Subject = "Test:Your in hand TR status"
	e.Text = []byte("Hi :\r\n This is a test mail.")
	e.HTML = []byte("Hi : \r\n<h1>Fancy Html is supported, too!</h1>")
	raw, err := e.Bytes()
	if err != nil {
		log.Fatal("Failed to render message: ", e)
	}
	m.writeFile(string(raw))
	return string(raw)
}

func (m *MailInfo) writeFile(content string) {
	fileName := CURRENT_DIR + "tmp.msg"
	flag, err := adutils.Exists(fileName)
	if flag {
		err = adutils.Unlink(fileName)
		if err != nil {
			os.Exit(1)
			return
		}
	}

	filefd, err := os.Create(fileName)
	if err != nil {
		os.Exit(1)
	}
	defer filefd.Close()
	m.ContentFile = fileName
	filefd.WriteString(content)

}
