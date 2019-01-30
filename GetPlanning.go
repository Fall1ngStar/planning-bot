package main

import (
	"io/ioutil"
	"net/http"
)

func GetWeeklyPlanning(yearMonthDayStart string, yearMonthDayEnd string) string {

	fileUrl := "http://adelb.univ-lyon1.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=9319&projectId=2&calType=ical&firstDate=" + yearMonthDayStart + "&lastDate=" + yearMonthDayEnd

	fileText := DownloadFileText(fileUrl)

	return fileText
}

func GetDailyPlanning(yearMonthDay string) string {

	fileUrl := "http://adelb.univ-lyon1.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=9319&projectId=2&calType=ical&firstDate=" + yearMonthDay + "&lastDate=" + yearMonthDay

	fileText := DownloadFileText(fileUrl)

	return fileText

}
func DownloadFileText(url string) string {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "Error : can't get http data"
	}
	defer resp.Body.Close()

	fileText, err := ioutil.ReadAll(resp.Body)

	println(string(fileText))

	return string(fileText)

}
