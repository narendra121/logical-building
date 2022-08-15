/*
calculation of age by using input from cmd
go run other/age-cal.go Y 2020-06-02
go run other/age-cal.go M 2020-06-02
go run other/age-cal.go D 2020-06-02
go run other/age-cal.go H 2020-06-02

*/
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"time"
)

const (
	lAYOUT = "2006-01-02"
)

func main() {
	inp := os.Args
	if len(os.Args) < 2 {
		log.Println("invalid input")
		return
	}
	caseEle := inp[1]
	dob := inp[2]
	ageCalculator(caseEle, dob)
}
func ageCalculator(caseEle, inp string) {
	ex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !ex.MatchString(inp) {
		log.Println("invalid dob format")
		return
	}
	db, err := time.Parse(lAYOUT, inp)
	if err != nil {
		log.Println("error in time.Parse() ", err)
		return
	}
	now := time.Now()
	res := now.Sub(db)
	switch caseEle {
	case "Y":
		yVal := ((res.Hours() / 24) / 30) / 12
		fmt.Println("Years: ", math.Round(yVal*100)/100)
	case "M":
		mVal := ((res.Hours() / 24) / 30)
		fmt.Println("Months: ", math.Round(mVal*100)/100)
	case "D":
		mVal := (res.Hours() / 24)
		fmt.Println("Days: ", math.Round(mVal*100)/100)
	case "H":
		mVal := (res.Hours())
		fmt.Println("Hours: ", math.Round(mVal*100)/100)
	default:
		fmt.Println("invalid data..")

	}

}
