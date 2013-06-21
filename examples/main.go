package main

import (
	"github.com/athom/gosegmentio"
)

func main() {
	gosegmentio.SetSecretToken("ep1ryb2az9awqg1xzy7i")
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
