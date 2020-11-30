package gohooks_test

import (
	"net/http"
	"testing"

	"github.com/averageflow/gohooks/v2/gohooks"
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

	resp, _ := hook.Send("http://www.google.com", map[string]string{"test": "test"})
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = http.MethodPatch

	resp, _ = hook.Send(receiverURL, map[string]string{"test": "test"})
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = http.MethodPut

	resp, _ = hook.Send(receiverURL, nil)
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = http.MethodDelete

	resp, _ = hook.Send(receiverURL, nil)
	if resp != nil {
		defer resp.Body.Close()
	}

	hook.PreferredMethod = "invalid"

	resp, _ = hook.Send(receiverURL, nil)
	if resp != nil {
		defer resp.Body.Close()
	}

	resp, _ = hook.Send("", nil)
	if resp != nil {
		defer resp.Body.Close()
	}

	resp, _ = hook.Send("ssh://google.com", nil)
	if resp != nil {
		defer resp.Body.Close()
	}
}

func TestCreateInvalidGoHook(t *testing.T) {
	hook := &gohooks.GoHook{}
	hook.Create(make(chan int), "", "")
}
