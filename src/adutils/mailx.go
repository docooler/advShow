package adutils

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func SendMail(to, content string) {
	// log.Println("mail :", to, "file: ", content)
	// args := content + " | sendmail" + to
	cmd := exec.Command("./sm.sh", content, to)
	_, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
	}

	if err := cmd.Start(); err != nil {
		log.Println(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err.Error())
	}
}

func SendMail_stub(to, subject, content string) {
	cmd := exec.Command("which", "ls")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err.Error())
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		log.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}

	log.Println("stdout:", string(bytes))
}
