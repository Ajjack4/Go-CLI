package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	dt := time.Now() 
	// cd:=dt.Format("01-02-2006 15:04:05")
    fmt.Println(dt.Format("01-02-2006 15:04:05")) 
	apiUrl :=fmt.Sprint("https://api.openf1.org/v1/sessions?year=2024")
	fmt.Println(apiUrl)
	res, err:=http.Get(apiUrl)
	if(err!=nil){
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Race Results not available")
	}
	body,err:=io.ReadAll(res.Body)
	if(err!=nil){
        panic(err)
    }
	fmt.Println(string(body))
	
 }
