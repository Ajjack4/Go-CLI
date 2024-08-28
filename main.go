package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	dt :=strconv.Itoa(time.Now().Year())
	// cd:=dt.Format("01-02-2006 15:04:05")
  
	apiUrl :=fmt.Sprintf("https://api.openf1.org/v1/sessions?year=%s", dt)
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
