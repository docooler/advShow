package adutils

import (
        "encoding/xml"
        "log"
        "io/ioutil"
        )

const (
        CONFIG_DIR   = "../../config/"
        SERVER_CFG   = "server.xml"
        CONTENT_CFG  = "content.xml"
      )

type Page struct {
    XMLName xml.Name `xml:"content"`
    Display []Display `xml:"dissplay"`
}

type Display struct {
    Name string `xml:"name"`
    Type string `xml:"type"`
    Link string `xml:"file"`
}


type Server struct {
     XMLName xml.Name `xml:"server"`
     Monitor `xml:"monitorsrv"`
     DisplayCtrl `xml:"displayctl"`
}

type Monitor struct {
    LocalServer string `xml:"localserver"`
}

type DisplayCtrl struct {
    Delaytime int `xml:"delaytime"`
    Transtime int `xml:"transtime"`
    Trmaxno   int `xml:"trmaxno"`
    Recmaxno  int `xml:"recmaxno"`
}


func ContentParse() (Page,error) {
    content , err := ioutil.ReadFile(CONFIG_DIR + CONTENT_CFG)
    var page Page
    if err != nil {
        log.Fatal(err)
        return page, err
    }
    
    err = xml.Unmarshal(content, &page)
    if err != nil {
        log.Fatal(err)
        log.Fatal("parser config failed")
        return page, err
    }
    
    return page , nil
}

func ServerParse()(Server, error){
    content , err := ioutil.ReadFile(CONFIG_DIR + SERVER_CFG)
    var page Server
    if err != nil {
        log.Fatal(err)
        return page, err
    }
    
    err = xml.Unmarshal(content, &page)
    if err != nil {
        log.Fatal(err)
        log.Fatal("parser config failed")
        return page, err
    }
    log.Println(page)
    return page , nil

}



