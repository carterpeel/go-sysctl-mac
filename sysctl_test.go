package tests

import (
	gosysctlmac "github.com/carterpeel/go-sysctl-mac"
	"testing"
)

func TestPersistent(t *testing.T) {
	// SetPersistent activates the config option now, but also adds it to
	// /private/etc/sysctl.conf so it persists over reboots.
	if err := gosysctlmac.SetPersistent("net.inet.ip.forwarding", "1"); err != nil {
		t.Fatalf("error enabling IP forwarding: %v\n", err)
	}
}

func TestNonPersistent(t *testing.T) {
	if err := sysctl.Set("kern.maxproc", "4096"); err != nil {
		t.Fatalf("error disabling IP forwarding: %v\n", err)
	}
}