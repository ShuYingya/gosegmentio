# GoSegmentIO

This Package is a simple solution for Segment.io API Go binding.
For the server side analysing.

[![Build Status](https://api.travis-ci.org/athom/gosegmentio.png?branch=master)](https://travis-ci.org/athom/gosegmentio)

## Installation

```
	go get "github.com/athom/gosegmentio"
```

## Usage

```go
package main

import (
	"github.com/athom/gosegmentio"
)

func main() {
	gosegmentio.SetSecretToken("ep1ryb2az9awqg1xzy7i")
	gosegmentio.Identify("athom", map[string]interface{}{
		"name":  "Athom",
		"email": "athom@126.com",
		"age":   26,
	}, map[string]interface{}{})

	gosegmentio.Track("007", "Avatar", map[string]interface{}{
		"Name":    "Kioshi",
		"Age":     303,
		"Country": "Earch Kindom",
	}, map[string]interface{}{})
}
```

## License

It is released under the [WTFPL License](http://www.wtfpl.net/txt/copying).
