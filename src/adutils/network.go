package adutils

import (
        "fmt"
        "net"
        "strings"
        )


func LocalIp() (string, error) {
    if addrs, err := net.InterfaceAddrs(); 
    err == nil {  
        for _, addr := range addrs {  
            fmt.Println(strings.Split(addr.String(),"/")[0])  
        }  
    }  
    return "123456", nil
}


func GetLocalIp() (string, error) {
    if addrs, err := net.InterfaceAddrs(); 
    err == nil {  
        for _, addr := range addrs {  
            fmt.Println(strings.Split(addr.String(),"/")[0]) 
            fmt.Println(strings.Split(addr.String(),"/")[1]) 
            fmt.Println(addr)
        }  
    }  
    return "123456", nil
}