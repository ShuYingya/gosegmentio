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
		"Country": "Earth Kindom",
	}, map[string]interface{}{})
}
