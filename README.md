### Gowoopra

[![Build Status](https://travis-ci.org/grsmv/gowoopra.svg)](https://travis-ci.org/grsmv/gowoopra)
[![Goreport](https://goreportcard.com/badge/github.com/grsmv/gowoopra)](https://goreportcard.com/report/github.com/grsmv/gowoopra)
[![godoc badge](http://godoc.org/github.com/grsmv/gowoopra?status.png)](http://godoc.org/github.com/grsmv/gowoopra)

Golang SDK for [Woopra](https://www.woopra.com) API based on  [API documentation](https://www.woopra.com/docs/developer/http-tracking-api/)

#### Usage:

```go
// defining new reusable Tracker with custom settings
wt, _ := gowoopra.NewTracker(map[string]string{

    // `host` is domain as registered in Woopra, it identifies which
    // project environment to receive the tracking request
    "host": "medcare.clinic",

    // In milliseconds, defaults to 30000 (equivalent to 30 seconds)
    // after which the event will expire and the visit will be marked
    // as offline.
    "timeout": 30000,
})

person := gowoopra.Person{
    Name: "Miles Davis",
    Email:"coltrane@johns.com",
}

// sending User-Agent HTTP header content as an optional argument
userAgent := r.UserAgent() // r is an idiomatic http.Request

// identifying current visitor in Woopra
id := wt.Identify(person, userAgent)

// Tracking custom event in Woopra. Each event can has additional data
id.Track(
    "login",           // event name
    map[string]string{ // custom data
        "through": "mobile",
        "when":    "yesterday",
        "mood":    "Really good",
    })

// it's possible to send only visitor's data to Woopra, without sending 
// any custom event and/or data
id.Push()
```

#### License:

```
Copyright (c) 2016 Serhii Herasymov

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```


