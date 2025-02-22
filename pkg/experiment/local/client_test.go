package local

import (
	"github.com/amplitude/experiment-go-server/pkg/experiment"
	"testing"
)

var client *Client

func init() {
	client = Initialize("server-qz35UwzJ5akieoAdIgzM4m9MIiOLXLoz", nil)
	err := client.Start()
	if err != nil {
		panic(err)
	}
}

func TestClientInitialize(t *testing.T) {
	client1 := Initialize("apiKey1", nil)
	client2 := Initialize("apiKey1", nil)
	client3 := Initialize("apiKey2", nil)
	if client1 != client2 {
		t.Fatalf("Expected equal client references.")
	}
	if client1 == client3 {
		t.Fatalf("Expected different client references.")
	}
}

func TestEvaluate(t *testing.T) {
	user := &experiment.User{UserId: "test_user"}
	result, err := client.Evaluate(user, nil)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	variant := result["sdk-local-evaluation-ci-test"]
	if variant.Key != "on" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "on" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Payload != "payload" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	variant = result["sdk-ci-test"]
	if variant.Key != "" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "" {
		t.Fatalf("Unexpected variant %v", variant)
	}
}


func TestEvaluateV2AllFlags(t *testing.T) {
	user := &experiment.User{UserId: "test_user"}
	result, err := client.EvaluateV2(user, nil)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	variant := result["sdk-local-evaluation-ci-test"]
	if variant.Key != "on" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "on" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Payload != "payload" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	variant = result["sdk-ci-test"]
	if variant.Key != "off" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "" {
		t.Fatalf("Unexpected variant %v", variant)
	}
}

func TestEvaluateV2OneFlag(t *testing.T) {
	user := &experiment.User{UserId: "test_user"}
	flagKeys := []string{"sdk-local-evaluation-ci-test"}
	result, err := client.EvaluateV2(user, flagKeys)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	variant := result["sdk-local-evaluation-ci-test"]
	if variant.Key != "on" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "on" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Payload != "payload" {
		t.Fatalf("Unexpected variant %v", variant)
	}
}

func TestEvaluateV2AllFlagsWithDependencies(t *testing.T) {
	user := &experiment.User{UserId: "user_id", DeviceId: "device_id"}
	result, err := client.EvaluateV2(user, nil)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	variant := result["sdk-ci-local-dependencies-test"]
	if variant.Key != "control" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "control" {
		t.Fatalf("Unexpected variant %v", variant)
	}
}

func TestEvaluateV2OneFlagWithDependencies(t *testing.T) {
	user := &experiment.User{UserId: "user_id", DeviceId: "device_id"}
	flagKeys := []string{"sdk-ci-local-dependencies-test"}
	result, err := client.EvaluateV2(user, flagKeys)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	variant := result["sdk-ci-local-dependencies-test"]
	if variant.Key != "control" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "control" {
		t.Fatalf("Unexpected variant %v", variant)
	}
}

func TestEvaluateV2UnknownFlagKey(t *testing.T) {
	user := &experiment.User{UserId: "user_id", DeviceId: "device_id"}
	flagKeys := []string{"does-not-exist"}
	result, err := client.EvaluateV2(user, flagKeys)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	variant := result["sdk-local-dependencies-test"]
	if variant.Key != "" {
		t.Fatalf("Unexpected variant %v", variant)
	}
	if variant.Value != "" {
		t.Fatalf("Unexpected variant %v", variant)
	}
}
