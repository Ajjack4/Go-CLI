package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type RaceSchedule struct {
    Races []Race 
}
type Race struct {
	name string `json:"country_name"`
	date string `json:"date_start"`
}

func main() {
	// currentDate := time.Now().Format("2006-01-02")
	dt :=strconv.Itoa(time.Now().Year())
	
	
	// fmt.Println(dt)
	// fmt.Println(currentDate)
	
   
	apiUrl :=fmt.Sprintf("https://api.openf1.org/v1/sessions?session_name=Race&year=%s",dt)
	// fmt.Println(apiUrl)
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
	// to maintain github strick ,if you are reading this get a life
 }
