package gowoopra

func Example() {

	// Sample package usage:
	wt, _ := NewTracker(map[string]string{
		"host": "sample-host.com",
	})

	wt.Identify(
		Person{"John Coltrane", "coltrane@johns.com"},
	).Track("login", map[string]string{
		"through": "mobile",
		"when":    "yesterday",
		"mood":    "Really good",
	})
}

func ExampleTracker_Identify() {

	// Usage with proxying client's `User-agent` to Woopra
	wt, _ := NewTracker(map[string]string{
		"host": "medcare.clinic",
	})

	wt.Identify(
		Person{"Miles Davis", "coltrane@johns.com"},

		// can be grabbed with http.Request.UserAgent()
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7",
	).Track("login",
		map[string]string{
			"through": "mobile",
			"when":    "yesterday",
			"mood":    "Really good",
		})
}

func ExampleContext_Push() {
	// defining new reusable Tracker with custom settings
	wt, _ := NewTracker(map[string]string{

		// `host` is domain as registered in Woopra, it identifies which
		// project environment to receive the tracking request
		"host": "medcare.clinic",

		// In milliseconds, defaults to 30000 (equivalent to 30 seconds)
		// after which the event will expire and the visit will be marked
		// as offline.
		"timeout": "30000",
	})

	person := Person{
		Name:  "Miles Davis",
		Email: "coltrane@johns.com",
	}

	// sending User-Agent HTTP header content as an optional argument
	userAgent := "Mozilla/5.0 (iPad; U; CPU [...]" // http.Request.UserAgent()

	// identifying current visitor in Woopra
	id := wt.Identify(person, userAgent)

	// Tracking custom event in Woopra. Each event can has additional data
	id.Track(
		"login", // event name
		map[string]string{ // custom data
			"through": "mobile",
			"when":    "yesterday",
			"mood":    "Really good",
		})

	// it's possible to send only visitor's data to Woopra, without sending
	// any custom event and/or data
	id.Push()
}
