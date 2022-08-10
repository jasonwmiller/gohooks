package gohooks_test

import (
	"encoding/json"
	"testing"

	"github.com/averageflow/gohooks/v2/gohooks"
)

func TestIsGoHookValid(t *testing.T) {
	secret := "0014716e-392c-4120-609e-555e295faff5" //nolint:gosec
	data := []int{1, 2, 3}

	hook := &gohooks.GoHook{}
	hook.Create(data, "int-resource", secret)

	type WebhookData struct {
		Resource string `json:"resource"`
		Data     []int  `json:"data"`
	}

	var hookData WebhookData

	err := json.Unmarshal(hook.PreparedData, &hookData)
	if err != nil {
		t.Errorf("Expected GoHook data to cleanly unmarshall to wanted JSON!")
	}

	isValid := gohooks.IsGoHookValid(hookData, hook.ResultingSha, secret)
	if !isValid {
		t.Errorf("Expected GoHook to be valid")
	}

	isValid = gohooks.IsGoHookValid(hookData, "invalid-signature", secret)
	if isValid {
		t.Errorf("Expected GoHook to be invalid")
	}

	isValid = gohooks.IsGoHookValid(make(chan int), "invalid-signature", "")
	if isValid {
		t.Errorf("Expected GoHook to be invalid")
	}
}
