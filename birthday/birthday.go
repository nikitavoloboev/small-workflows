package main

import (
	"log"
	"os"
	"time"

	aw "github.com/deanishe/awgo"
	humanize "github.com/dustin/go-humanize"
)

var (
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}

	dateOfBirth = os.Getenv("Date of Birth (DD/MM/YYYY)")

	query string
)

// Gets age correctly.
func getAge(birthday time.Time) int {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()

	birthYear := birthday.Year()
	birthMonth := birthday.Month()
	birthDay := birthday.Day()

	var age int
	if currentMonth < birthMonth {
		age = currentYear - birthYear - 1
	} else if currentMonth > birthMonth {
		age = currentYear - birthYear
	} else {
		if currentDay >= birthDay {
			age = currentYear - birthYear
		} else {
			age = currentYear - birthYear - 1
		}
	}
	return age
}

// TODO: cache value so not to rerun this every second, only append to seconds on reruns
func getBirthday() error {
	log.Printf("query=%s", query)

	// Get birthday
	if dateOfBirth == "" {
		wf.NewItem("Please add your date of birth in configuration sheet")
		wf.SendFeedback()
	} else {
		wf.Rerun(1)
		const dateForm = "02/01/2006"
		parsedDate, err := time.Parse(dateForm, dateOfBirth)
		if err != nil {
			log.Fatal(err)
		}

		birthDate := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.UTC)

		timeSinceBirthdate := time.Since(birthDate)

		wf.NewItem("You have lived:")
		// Years
		wf.NewItem(humanize.Comma(int64(timeSinceBirthdate.Hours()/8760)) + " years")
		// Months
		wf.NewItem(humanize.Comma(int64(timeSinceBirthdate.Hours()/730)) + " months")
		// Days
		wf.NewItem(humanize.Comma(int64(timeSinceBirthdate.Hours()/24)) + " days")
		// Hours
		wf.NewItem(humanize.Comma(int64(timeSinceBirthdate.Hours())) + " hours")
		// Minutes
		wf.NewItem(humanize.Comma(int64(timeSinceBirthdate.Minutes())) + " minutes")
		// Seconds
		wf.NewItem(humanize.Comma(int64(timeSinceBirthdate.Seconds())) + " seconds")

		query = os.Args[1]

		if query != "" {
			wf.Filter(query)
		}

		wf.WarnEmpty("No matching items", "Try a different query?")
		wf.SendFeedback()
	}

	return nil
}
