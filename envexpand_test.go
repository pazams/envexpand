package envexpand_test

import (
	"bytes"
	"github.com/pazams/envexpand"
	"os"
	"testing"
)

func TestExpands(t *testing.T) {

	os.Setenv("ENVEXPAND_DOMAIN", "google.com")
	defer os.Unsetenv("ENVEXPAND_DOMAIN")

	input := []byte("http://${ENVEXPAND_DOMAIN}")
	expected := []byte("http://google.com")
	output := envexpand.Expand(input)
	if !bytes.Equal(output, expected) {
		t.Errorf("output does not match expected", input, output, expected)
	}
}

func TestExpandsString(t *testing.T) {

	os.Setenv("ENVEXPAND_DOMAIN", "google.com")
	defer os.Unsetenv("ENVEXPAND_DOMAIN")

	input := "http://${ENVEXPAND_DOMAIN}"
	expected := "http://google.com"
	output := envexpand.ExpandString(input)

	if output != expected {
		t.Errorf("output does not match expected", input, output, expected)
	}
}

func TestDoesNotExpandIfNotFound(t *testing.T) {

	input := []byte("http://${ENVEXPAND_DOMAIN}")
	expected := []byte("http://${ENVEXPAND_DOMAIN}")
	output := envexpand.Expand(input)

	if !bytes.Equal(output, expected) {
		t.Errorf("output does not match expected", input, output, expected)
	}
}

func TestDoesNotExpandIfNotFoundString(t *testing.T) {

	input := "http://${ENVEXPAND_DOMAIN}"
	expected := "http://${ENVEXPAND_DOMAIN}"
	output := envexpand.ExpandString(input)

	if output != expected {
		t.Errorf("output does not match expected", input, output, expected)
	}
}

func TestExpandsMultiple(t *testing.T) {

	os.Setenv("ENVEXPAND_DOMAIN", "google.com")
	defer os.Unsetenv("ENVEXPAND_DOMAIN")
	os.Setenv("ENVEXPAND_PORT", "8080")
	defer os.Unsetenv("ENVEXPAND_PORT")

	input := []byte("http://${ENVEXPAND_DOMAIN}:${ENVEXPAND_PORT}")
	expected := []byte("http://google.com:8080")
	output := envexpand.Expand(input)

	if !bytes.Equal(output, expected) {
		t.Errorf("output does not match expected", input, output, expected)
	}
}

func TestExpandsMultipleString(t *testing.T) {

	os.Setenv("ENVEXPAND_DOMAIN", "google.com")
	defer os.Unsetenv("ENVEXPAND_DOMAIN")
	os.Setenv("ENVEXPAND_PORT", "8080")
	defer os.Unsetenv("ENVEXPAND_PORT")

	input := "http://${ENVEXPAND_DOMAIN}:${ENVEXPAND_PORT}"
	expected := "http://google.com:8080"
	output := envexpand.ExpandString(input)

	if output != expected {
		t.Errorf("output does not match expected", input, output, expected)
	}
}
