package main

import (
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    "github.com/laurent22/ical-go"
    "io/ioutil"
    "net/http"
    "os"
    "sort"
    "strconv"
    "time"
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

    return string(fileText)

}

type Lesson struct {
    Summary  string
    Location string
    date     time.Time
}

func GetLessons(data string) []*Lesson {
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

            lessons = append(lessons, &Lesson{
                date:     t,
                Summary:  summary,
                Location: location,
            })
        }
    }

    sort.Slice(lessons, func(i, j int) bool {
        return lessons[i].date.Before(lessons[j].date)
    })
    return lessons
}

func tweet(message string) {
    config := oauth1.NewConfig(os.Getenv("BOT_KEY"), os.Getenv("BOT_KEY_SECRET"))
    token := oauth1.NewToken(os.Getenv("BOT_TOKEN"), os.Getenv("BOT_TOKEN_SECRET"))
    httpClient := config.Client(oauth1.NoContext, token)

    // Twitter client
    client := twitter.NewClient(httpClient)

    // Send a Tweet
    _, _, err := client.Statuses.Update(message, nil)
    if err != nil {
        fmt.Println(err)
    }

}

func handler() {
    now := time.Now()
    planning := GetDailyPlanning(strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day()))
    lessons := GetLessons(planning)
    for _, lesson := range lessons {
        if lesson.date.Hour() > now.Hour() {
            message := "Next lesson: " + lesson.Summary + "\nLocation: " + lesson.Location + "\nHour: " + strconv.Itoa(lesson.date.Hour() + 1) + ":" + strconv.Itoa(lesson.date.Minute())
            tweet(message)
            break
        }
    }
}

func main() {
    lambda.Start(handler)
}
