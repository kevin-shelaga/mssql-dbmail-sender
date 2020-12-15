package helpers

import (
	"testing"
)

func TestGetDateTime(t *testing.T) {
	result := GetDateTime()

	if result == "" {
		t.Errorf("GetDateTime() = \"\"; want time.Now formatted")
	}
}

func BenchmarkGetDateTime(t *testing.B) {

	GetDateTime()
}

func TestEncodeRFC2047(t *testing.T) {
	result := encodeRFC2047("test string")

	if result == "" {
		t.Errorf("test failed!")
	}
}

func BenchmarkEncodeRFC2047(t *testing.B) {
	result := encodeRFC2047("test string")

	if result == "" {
		t.Errorf("test failed!")
	}
}

func TestComposeMimeMail(t *testing.T) {
	result := ComposeMimeMail("to", "from", "subject", "body", "importance", "format")

	if result == nil {
		t.Errorf("test failed!")
	}
}

func BenchmarkComposeMimeMail(t *testing.B) {
	result := ComposeMimeMail("to", "from", "subject", "body", "importance", "format")

	if result == nil {
		t.Errorf("test failed!")
	}
}

func TestGetConfig(t *testing.T) {
	result := GetConfig("../appsettings.json")

	if result == nil {
		t.Errorf("test failed!")
	}
}

func BenchmarkGetConfig(t *testing.B) {
	result := GetConfig("../appsettings.json")

	if result == nil {
		t.Errorf("test failed!")
	}
}