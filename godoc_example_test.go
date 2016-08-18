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