package calendarservice

import (
	"context"
	"strings"

	calendarpb "github.com/benkim0414/calendarservice/pb"
	"github.com/golang/protobuf/ptypes/empty"
	calendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type service struct {
	client *calendar.Service
}

func New(ctx context.Context, opts ...option.ClientOption) (*service, error) {
	client, err := calendar.NewService(ctx, option.WithScopes(calendar.CalendarEventsScope))
	if err != nil {
		return nil, err
	}
	return &service{client: client}, nil
}

func (svc *service) ListEvents(ctx context.Context, req *calendarpb.ListEventsRequest) (*calendarpb.ListEventsResponse, error) {
	calendarID := strings.Split(req.Parent, "/")[1]
	res, err := svc.client.Events.List(calendarID).Do()
	if err != nil {
		return &calendarpb.ListEventsResponse{}, err
	}
	events := make([]*calendarpb.Event, 0)
	for _, event := range res.Items {
		events = append(events, newEvent(event))
	}
	return &calendarpb.ListEventsResponse{
		Events:        events,
		NextPageToken: res.NextPageToken,
	}, nil
}

func (svc *service) GetEvent(ctx context.Context, req *calendarpb.GetEventRequest) (*calendarpb.Event, error) {
	ss := strings.Split(req.Name, "/")
	calendarID, eventID := ss[1], ss[3]
	event, err := svc.client.Events.Get(calendarID, eventID).Do()
	if err != nil {
		return &calendarpb.Event{}, err
	}
	return newEvent(event), nil
}

func (svc *service) CreateEvent(ctx context.Context, req *calendarpb.CreateEventRequest) (*calendarpb.Event, error) {
	calendarID := strings.Split(req.Parent, "/")[1]
	event, err := svc.client.Events.QuickAdd(calendarID, req.Text).Do()
	if err != nil {
		return &calendarpb.Event{}, err
	}
	return newEvent(event), nil
}

func (svc *service) DeleteEvent(ctx context.Context, req *calendarpb.DeleteEventRequest) (*empty.Empty, error) {
	ss := strings.Split(req.Name, "/")
	calendarID, eventID := ss[1], ss[3]
	err := svc.client.Events.Delete(calendarID, eventID).Do()
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}

func newEvent(event *calendar.Event) *calendarpb.Event {
	attendees := make([]*calendarpb.Attendee, 0)
	for _, attendee := range event.Attendees {
		attendees = append(attendees, newAttendee(attendee))
	}
	return &calendarpb.Event{
		Id:          event.Id,
		Status:      event.Status,
		HtmlLink:    event.HtmlLink,
		CreateTime:  event.Created,
		UpdateTime:  event.Updated,
		Summary:     event.Summary,
		Description: event.Description,
		Location:    event.Location,
		Creator:     event.Creator.Email,
		Organizer:   event.Organizer.Email,
		StartTime:   event.Start.DateTime,
		EndTime:     event.End.DateTime,
		ICalUid:     event.ICalUID,
		Attendees:   attendees,
	}
}

func newAttendee(attendee *calendar.EventAttendee) *calendarpb.Attendee {
	return &calendarpb.Attendee{
		Id:               attendee.Id,
		Email:            attendee.Email,
		DisplayName:      attendee.DisplayName,
		Organizer:        attendee.Organizer,
		Optional:         attendee.Optional,
		ResponseStatus:   attendee.ResponseStatus,
		Comment:          attendee.Comment,
		AdditionalGuests: attendee.AdditionalGuests,
	}
}
