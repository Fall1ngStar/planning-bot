package main

import (
	"fmt"
	"github.com/laurent22/ical-go"
	"time"
)

type Lesson struct {
	Summary  string
	Location string
	date     time.Time
}

func main() {

}

func getLessons(data string) []*Lesson {
	parsed, _ := ical.ParseCalendar(data)

	lessons := make([]*Lesson, 0)
	for _, value := range parsed.Children {

		summary := value.PropString("SUMMARY", "No summary")
		startDate := value.PropString("DTSTART", "No start date")
		location := value.PropString("LOCATION", "No location")


		if startDate != "No start date" {
			var t time.Time

			year := startDate[0:4]
			month := startDate[4:6]
			day := startDate[6:8]
			hours := startDate[9:11]
			minutes := startDate[11:13]
			seconds := startDate[13:15]
			date := year + "-" + month + "-" + day + "T" + hours + ":" + minutes + ":" + seconds + ".000Z"

			layout := "2006-01-02T15:04:05.000Z"
			t, err := time.Parse(layout, date)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t)

			//fmt.Println(date)

			//fmt.Println(summary)
			//fmt.Println(startDate)
			//fmt.Println(location)


			lessons = append(lessons, &Lesson{
				date: t,
				Summary:summary,
				Location:location,
			})
		}
	}

	return lessons
}