package gohooks_test

import (
	"net/http"
	"testing"

	"github.com/jasonwmiller/gohooks/v2/gohooks"
)

const (
	receiverURL = "http://www.google.com"
)

func TestSend(t *testing.T) {
	secret := "2014716e-392c-4120-609e-555e295faff5" //nolint:gosec
	data := []int{1, 2, 3}

	hook := &gohooks.GoHook{}
	hook.CreateWithoutWrapper(data, secret)

	hook = &gohooks.GoHook{}
	hook.Create(data, "int-resource", secret)

	resp, _ := hook.Send("http://www.google.com")
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = http.MethodPatch

	resp, _ = hook.Send(receiverURL)
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = http.MethodPut

	resp, _ = hook.Send(receiverURL)
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = http.MethodDelete
	hook.AdditionalHeaders = map[string]string{"test": "test"}

	resp, _ = hook.Send(receiverURL)
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = "invalid"

	resp, _ = hook.Send(receiverURL)
	if resp != nil {
		defer resp.Body.Close()
	}

	resp, _ = hook.Send("")
	if resp != nil {
		defer resp.Body.Close()
	}

	resp, _ = hook.Send("ssh://google.com")
	if resp != nil {
		defer resp.Body.Close()
	}
}

func TestCreateInvalidGoHook(t *testing.T) {
	hook := &gohooks.GoHook{}
	hook.Create(make(chan int), "", "")
}
