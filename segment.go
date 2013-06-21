package gosegmentio

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	ErrTimeOut = errors.New("Request Time Out")
	ErrUnknown = errors.New("Unknow Error")
)

const (
	BASE_URL = "https://api.segment.io/v1/"
)

var DefaultPusher = NewPusher("")
var DefaultTimeOut = 3 * time.Second

func SetTimeOut(to time.Duration) {
	DefaultTimeOut = to
}

func SetSecretToken(token string) {
	DefaultPusher.secretToken = token
}

func NewPusher(token string) (r *Pusher) {
	return &Pusher{
		secretToken: token,
	}
}

type Pusher struct {
	secretToken string
}

type ResultError struct {
	Type    string `json"type"`
	Message string `json"message"`
}

type ResultJson struct {
	Success bool `json:"success"`
	Error   *ResultError
}

func (this *Pusher) Do(verb string, m interface{}, ech chan error) {
	data, err := json.Marshal(m)
	if err != nil {
		ech <- err
		return
	}

	resp, err := http.Post(BASE_URL+verb, "application/json", strings.NewReader(string(data)))
	if err != nil {
		ech <- err
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ech <- err
		return
	}
	rj := &ResultJson{}
	json.Unmarshal(body, &rj)
	if !rj.Success {
		if rj.Error != nil {
			err = errors.New(rj.Error.Type)
		} else {
			err = ErrUnknown
		}
		ech <- err
		return
	}
	ech <- nil
	return
}

func (this *Pusher) SendData(verb string, m interface{}) (err error) {
	ech := make(chan error)
	go this.Do(verb, m, ech)
	select {
	case err = <-ech:
		return
	case <-time.After(DefaultTimeOut):
		err = ErrTimeOut
		return
	}
	return
}

func Identify(userId string, traits map[string]interface{}, context map[string]interface{}) (err error) {
	return DefaultPusher.Identify(userId, traits, context)
}

func (this *Pusher) Identify(userId string, traits map[string]interface{}, context map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["secret"] = this.secretToken
	m["userId"] = userId
	p := map[string]interface{}{}
	for k, v := range traits {
		p[k] = v
	}
	m["traits"] = p
	c := map[string]interface{}{}
	for k, v := range context {
		c[k] = v
	}
	m["context"] = c
	return this.SendData("identify", m)
}

func Track(userId string, eventName string, eventProperties map[string]interface{}, options map[string]interface{}) (err error) {
	return DefaultPusher.Track(userId, eventName, eventProperties, options)
}

func (this *Pusher) Track(userId string, eventName string, eventProperties map[string]interface{}, options map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["secret"] = this.secretToken
	m["userId"] = userId
	m["event"] = eventName
	p := map[string]interface{}{}
	for k, v := range eventProperties {
		p[k] = v
	}
	m["properties"] = p
	c := map[string]interface{}{}
	for k, v := range options {
		c[k] = v
	}
	m["context"] = c
	return this.SendData("track", m)
}

func Alias(fromId string, toId string, context map[string]interface{}) (err error) {
	return DefaultPusher.Alias(fromId, toId, context)
}

func (this *Pusher) Alias(fromId string, toId string, context map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["secret"] = this.secretToken
	m["from"] = fromId
	m["to"] = toId
	c := map[string]interface{}{}
	for k, v := range context {
		c[k] = v
	}
	m["context"] = c
	return this.SendData("alias", m)
}
