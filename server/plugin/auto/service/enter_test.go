package service

import "testing"

func TestServiceGroupAppProvidesAutoServices(t *testing.T) {
	if ServiceGroupApp == nil {
		t.Fatalf("expected plugin service group to be initialized")
	}
}
