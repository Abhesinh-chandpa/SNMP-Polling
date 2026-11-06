package app

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	var out bytes.Buffer
	if err := Run([]string{"help"}, &out); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "monitor") {
		t.Fatalf("missing roadmap command: %s", out.String())
	}
}

func TestVersion(t *testing.T) {
	var out bytes.Buffer
	if err := Run([]string{"version"}, &out); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), Version) {
		t.Fatal(out.String())
	}
}

func TestUnknownCommand(t *testing.T) {
	var out bytes.Buffer
	if err := Run([]string{"unexpected"}, &out); err == nil {
		t.Fatal("expected error")
	}
}
