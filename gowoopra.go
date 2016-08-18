package gowoopra

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

const (
	TRACK_API_ENDPOINT    = "https://www.woopra.com/track/ce"
	IDENTIFY_API_ENDPOINT = "https://www.woopra.com/track/identify"
	NO_HOST_ERROR         = "Host should be specified. This is domain as registered in Woopra, it identifies which project environment to receive the tracking request"
)

// Tracker used for storing sharable settings
type Tracker struct {
	Host    string
	Timeout int
}

// Person used for identifying person (user) for upcoming tracking
type Person struct {
	Name  string
	Email string
}

// Context is an intermediate representation of concatenated track-request specific data
type Context struct {
	Tracker
	Person
	Event             string
	VisitorProperties map[string]string
	UserAgent         string
}

// NewTracker creates new instance of sharable Tracker struct
func NewTracker(config map[string]string) (*Tracker, error) {
	if len(config["host"]) == 0 {
		return &Tracker{}, errors.New(NO_HOST_ERROR)
	}
	return &Tracker{Host: config["host"]}, nil
}

// Identify is used to create reusable person-specific association.
// Last (optional) argument is used to proxy UserAgent data to Woopra
func (t Tracker) Identify(person Person, args ...string) *Context {
	var userAgent string
	if len(args) > 0 {
		userAgent = args[0]
	}
	return &Context{
		Tracker:   t,
		Person:    person,
		UserAgent: userAgent,
	}
}

// Push identify a user without any tracking event
func (ctx *Context) Push() (*Context) {
	ctx.performRequest(IDENTIFY_API_ENDPOINT)
	return ctx
}

// Track is tracking a custom event for identified person
func (ctx *Context) Track(event string, properties map[string]string) *Context {
	ctx.Event = event
	ctx.VisitorProperties = properties
	ctx.performRequest(TRACK_API_ENDPOINT)
	return ctx
}

// performRequest is a low-level abstraction, used for sending requests to the Woopra API
func (ctx *Context) performRequest(endpoint string) {
	req, _ := http.NewRequest("POST", endpoint, nil)

	// proxying User-Agent header from client to Woopra
	if len(ctx.UserAgent) > 0 {
		header := make(http.Header)
		header.Add("User-Agent", ctx.UserAgent)
		req.Header = header
	}

	req.URL.RawQuery = ctx.prepareQuery()

	go func() {
		resp, _ := http.DefaultClient.Do(req)
		defer resp.Body.Close()
	}()
}

// prepareQuery concatenates all needed request parameters
func (ctx *Context) prepareQuery() string {
	var values = make(url.Values)
	values.Add("host", ctx.Tracker.Host)

	if ctx.Tracker.Timeout > 0 {
		values.Add("timeout", strconv.Itoa(ctx.Tracker.Timeout))
	}

	if len(ctx.Person.Email) > 0 {
		values.Add("cv_email", ctx.Person.Email)
	}

	if len(ctx.Person.Name) > 0 {
		values.Add("cv_name", ctx.Person.Name)
	}

	if len(ctx.Event) > 0 {
		values.Add("event", ctx.Event)
	}

	for k, v := range ctx.VisitorProperties {
		values.Add("ce_"+k, v)
	}
	return values.Encode()
}
