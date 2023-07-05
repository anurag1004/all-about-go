package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	Msg      string
	LoggedAt string
}

func (log Log) getLog() string {
	return fmt.Sprintf("%v : %v", log.LoggedAt, log.Msg)
}
func (log Log) getUnixTime() int64 {
	layout := "01/02/2006 15:04:05"
	datetime, err := time.Parse(layout, log.LoggedAt)
	if err != nil {
		panic(err)
	}
	unixTimestamp := datetime.Unix()
	return unixTimestamp
}

// func (log Log) MarshalJSON() ([]byte, error) {
// 	customStruct := struct {
// 		LoggedAt int64
// 		Log
// 	}{
// 		LoggedAt: log.getUnixTime(),
// 		Log:      (log),
// 	}
// 	return json.Marshal(customStruct)
// }

func (log Log) MarshalJSON() ([]byte, error) {
	type Allias Log
	customStruct := struct {
		LoggedAt int64
		Allias
	}{
		LoggedAt: log.getUnixTime(),
		Allias:   (Allias)(log),
	}
	return json.Marshal(customStruct)
}
func main() {
	log := Log{
		Msg:      "Sample Log",
		LoggedAt: "12/11/2023 11:09:49",
	}
	bs, _ := json.Marshal(log)
	fmt.Println(string(bs))

}
