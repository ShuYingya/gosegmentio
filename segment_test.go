package gosegmentio

import (
	"testing"
	"time"
)

func TestTrack(t *testing.T) {
	var err error
	SetSecretToken("ep1ryb2az9awqg1xzy7i")
	err = Track("athom", "Profile", map[string]interface{}{
		"name":  "Athom",
		"email": "athom@126.com",
		"age":   26,
	}, map[string]interface{}{})
	if err != nil {
		t.Errorf("real token should get no error")
	}

	SetSecretToken("i like to move it move it n mm")
	err = Track("athom", "Profile", map[string]interface{}{
		"name":   "Athom",
		"email":  "athom@126.com",
		"age":    26,
		"gender": "m",
	}, map[string]interface{}{})
	if err == nil {
		t.Errorf("fake token, should get error:" + err.Error())
	}

	SetTimeOut(10 * time.Millisecond)
	err = Track("athom", "Profile", map[string]interface{}{
		"name":   "Athom",
		"email":  "athom@126.com",
		"age":    26,
		"gender": "m",
	}, map[string]interface{}{})
	if err == nil {
		t.Errorf("time is up, should get error:" + err.Error())
	}
	if err != ErrTimeOut {
		t.Errorf("should get timeout error")
	}
}
