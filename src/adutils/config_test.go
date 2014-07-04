package adutils

import (
        "testing"
        "log"
        )

func Test_ContentParse(t *testing.T) {
    page,err  := ContentParse()
    if err != nil {
        panic(err)
        return
    }
    log.Println(page)
    log.Println(page.Display[0].Name)
    log.Println(len(page.Display))
}    

func Test_ServerParse(t *testing.T){
    _,err  := ServerParse()
    if err != nil {
        panic(err)
    }
}