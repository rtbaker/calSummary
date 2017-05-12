package main

import (
//	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
//	"net/http"
//	"net/url"
//	"os"
	"os/user"
	"time"
	"strings"
	"golang.org/x/net/context"
//	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"path/filepath"
)


func main() {
	ctx := context.Background()

	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	serviceCredsFile := filepath.Join(usr.HomeDir, ".CalProject.json")

	b, err := ioutil.ReadFile(serviceCredsFile)
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/calendar-go-quickstart.json
	config, err := google.JWTConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	//client := getClient(ctx, config)
	client := config.Client(ctx)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}

	now := time.Now()
	then := now.AddDate(0,1,0)

	tMin := now.Format(time.RFC3339)
	tMax := then.Format(time.RFC3339)

	events, err := srv.Events.List("e5dtcl0a60cnlu00ma9mses6sk@group.calendar.google.com").ShowDeleted(false).
		SingleEvents(true).TimeMin(tMin).TimeMax(tMax).MaxResults(99).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	// fmt.Println("Upcoming events:")
	myEvents := make([]string, 0, 0)

	if len(events.Items) > 0 {
		for _, i := range events.Items {
			var when string
			var to string
			to = ""

			// If the DateTime is an empty string the Event is an all-day Event.
			// So only Date is available.
			if i.Start.DateTime != "" {
				when = i.Start.DateTime
			} else {
				when = i.Start.Date
				to = i.End.Date
			}

			if len(to) != 0{
				whenT, _ := time.Parse("2006-01-02", when)
				toT, _ := time.Parse("2006-01-02", to)

				str := fmt.Sprintf("%s (%s - %s)", i.Summary, whenT.Format("02/01/06"), toT.Format("02/01/06"))
				myEvents = append(myEvents, str)
			} else {
				whenT, _ := time.Parse(time.RFC3339, when)

				str := fmt.Sprintf("%s (%s)", i.Summary, whenT.Format("02/01/06 15:04"))
				myEvents = append(myEvents, str)
			}

		}

		fmt.Println(strings.Join(myEvents, " *** "))
	} else {
		fmt.Printf("No upcoming events found.\n")
	}

}
