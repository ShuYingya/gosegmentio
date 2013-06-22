# GoSegmentIO

This Package is a simple solution for Segment.io API Go binding.
For the server side analysing.

[![Build Status](https://api.travis-ci.org/athom/gosegmentio.png?branch=master)](https://travis-ci.org/athom/gosegmentio)

## Installation

```
	go get "github.com/athom/gosegmentio"
```

## Usage

- #### Simple Usage:

```go
package main

import (
	"github.com/athom/gosegmentio"
)

func main() {
	gosegmentio.SetSecretToken("your secrect token")
	gosegmentio.Identify("112", map[string]interface{}{
		"Name":    "Kioshi",
		"Age":     230,
		"Country": "Earth Kindom",
	}, map[string]interface{}{})

	gosegmentio.Track("112", "DeadAt", map[string]interface{}{
		"Age":      230,
		"Date":     "82 BSC",
		"Location": "Kioshi Island",
	}, map[string]interface{}{})

	gosegmentio.Track("102", "AnomynousEvent", map[string]interface{}{
		"Hey": "There",
	}, map[string]interface{}{})

	gosegmentio.Alias("102", "113", map[string]interface{}{})

	gosegmentio.Identify("113", map[string]interface{}{
		"Name":    "Roku",
		"Age":     70,
		"Country": "Fire Nation",
	}, map[string]interface{}{})

	gosegmentio.Track("113", "BornAt", map[string]interface{}{
		"Age":      0,
		"Date":     "82 BSC",
		"Location": "Fire Nation",
	}, map[string]interface{}{})

	gosegmentio.Track("113", "DeadAt", map[string]interface{}{
		"Age":      70,
		"Date":     "12 BSC",
		"Location": "Fire Nation",
	}, map[string]interface{}{})
}
```

- #### Advanced(the batch feature):

Please refer the [examples](https://github.com/athom/gosegmentio/blob/master/examples/batch_demo.go)

## License

It is released under the [WTFPL License](http://www.wtfpl.net/txt/copying).
