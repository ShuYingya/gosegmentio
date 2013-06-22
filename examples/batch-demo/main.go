package main

import (
	"fmt"
	"github.com/athom/gosegmentio"
)

func main() {
	gosegmentio.SetSecretToken("ep1ryb2az9awqg1xzy7i")
	identifyKioshi := map[string]interface{}{
		"action": "identify",
		"userId": "yy_112",
		"traits": map[string]interface{}{
			"Name":    "Kioshi",
			"Age":     230,
			"Country": "Earth Kindom",
		},
	}
	trackKioshi := map[string]interface{}{
		"action": "track",
		"userId": "yy_112",
		"event":  "DeadAt",
		"properties": map[string]interface{}{
			"Age":      230,
			"Date":     "82 BSC",
			"Location": "Kioshi Island",
		},
	}
	trackAnonym := map[string]interface{}{
		"action": "track",
		"userId": "yy_102",
		"event":  "Anonymous Shouting",
		"properties": map[string]interface{}{
			"Hey": "There",
		},
	}
	aliasAnonymousRoku := map[string]interface{}{
		"action": "alias",
		"from":   "yy_102",
		"to":     "yy_113",
	}
	identifyRoku := map[string]interface{}{
		"action": "identify",
		"userId": "yy_113",
		"traits": map[string]interface{}{
			"Name":    "Roku",
			"Age":     70,
			"Country": "Fire Nation",
		},
	}
	trackRokuBirth := map[string]interface{}{
		"action": "track",
		"userId": "yy_113",
		"event":  "BornAt",
		"properties": map[string]interface{}{
			"Age":      0,
			"Date":     "82 BSC",
			"Location": "Fire Nation",
		},
	}
	trackRokuDeath := map[string]interface{}{
		"action": "track",
		"userId": "yy_113",
		"event":  "DeadAt",
		"properties": map[string]interface{}{
			"Age":      0,
			"Date":     "82 BSC",
			"Location": "Fire Nation",
		},
	}

	batch := []map[string]interface{}{
		identifyKioshi,
		trackKioshi,
		trackAnonym,
		aliasAnonymousRoku,
		identifyRoku,
		trackRokuBirth,
		trackRokuDeath,
	}

	err := gosegmentio.Import(batch)
	if err != nil {
		fmt.Println(err)
	}
}
