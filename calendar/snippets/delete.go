package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func deleteEvent(client *http.Client) {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// specify the calendar id for which events will be created
	calendarId := "primary"
	// Retrieve the event from the API
	eventID := "id of the created event"
	err = srv.Events.Delete(calendarId, eventID).SendUpdates("all").Do()
	if err != nil {
		log.Fatalf("Unable to delete calendar event %v", err)
	}

	fmt.Printf("Event deleted")
}