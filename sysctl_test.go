package sysctl

import (
	"testing"
)

func TestPersistent(t *testing.T) {
	t.Cleanup(func() {
		if err := Set("net.inet.ip.forwarding", "0"); err != nil {
			t.Fatalf("error disabling IP forwarding: %v\n", err)
		}
	})

	// SetPersistent() activates the config option now, but also adds it to
	// /private/etc/sysctl.conf, so it persists over reboots.
	if err := SetPersistent("net.inet.ip.forwarding", "1"); err != nil {
		t.Fatalf("error enabling IP forwarding: %v\n", err)
	}
}

func TestNonPersistent(t *testing.T) {
	t.Cleanup(func() {
		if err := Set("kern.maxproc", "2048"); err != nil {
			t.Fatalf("error resetting maxprocs to default: %v\n", err)
		}
	})

	// Set() sets a config option that will not persist upon reboot.
	if err := Set("kern.maxproc", "4096"); err != nil {
		t.Fatalf("error incrementing maxproc: %v\n", err)
	}
}