package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err:=http.Get("https://api.openf1.org/v1/position?meeting_key=1217&driver_number=40&position<=3")
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
