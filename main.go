package main

import (
//	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
//	"net/http"
//	"net/url"
//	"os"
//	"os/user"
//	"path/filepath"
	"time"
	"golang.org/x/net/context"
//	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)


func main() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("./CalProject.json")
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

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("e5dtcl0a60cnlu00ma9mses6sk@group.calendar.google.com").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	fmt.Println("Upcoming events:")
	if len(events.Items) > 0 {
		for _, i := range events.Items {
			var when string
			// If the DateTime is an empty string the Event is an all-day Event.
			// So only Date is available.
			if i.Start.DateTime != "" {
				when = i.Start.DateTime
			} else {
				when = i.Start.Date
			}
			fmt.Printf("%s (%s)\n", i.Summary, when)
		}
	} else {
		fmt.Printf("No upcoming events found.\n")
	}

}
