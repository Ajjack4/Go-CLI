package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Ajjack4/Go-CLI/Api"
)

type RaceSchedule []Race 

type Race struct {
	Name string `json:"country_name"`
	Date string `json:"date_start"`
}

func main() {
    currentTime:=time.Now()
	formattedtime:=currentTime.Format("2006")
	
   Api.Apicall()
	apiUrl :=fmt.Sprintf("https://api.openf1.org/v1/sessions?session_name=Race&year=%s",formattedtime)
  
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
	
	var schedule RaceSchedule
	err = json.Unmarshal(body, &schedule)
	if err!= nil {
        panic(err)
    }
	fmt.Println(schedule)
	
 }
