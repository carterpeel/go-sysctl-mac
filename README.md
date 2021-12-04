# go-sysctl-mac

Allows you to programatically modify key:value sysctl pairs on MacOS.

Tested & confirmd to work on MacOS Big Sur & newer releases.

## Example:
```go
package mypkg

import (
  "testing"
  "github.com/carterpeel/go-sysctl-mac"
)

func TestPersistent(t *testing.T) {
	t.Cleanup(func() {
		if err := sysctl.Set("net.inet.ip.forwarding", "0"); err != nil {
			t.Fatalf("error disabling IP forwarding: %v\n", err)
		}
	})

	// SetPersistent() activates the config option now, but also adds it to
	// /private/etc/sysctl.conf, so it persists over reboots.
	if err := sysctl.SetPersistent("net.inet.ip.forwarding", "1"); err != nil {
		t.Fatalf("error enabling IP forwarding: %v\n", err)
	}
}

func TestNonPersistent(t *testing.T) {
	t.Cleanup(func() {
		if err := sysctl.Set("kern.maxproc", "2048"); err != nil {
			t.Fatalf("error resetting maxprocs to default: %v\n", err)
		}
	})

	// Set() sets a config option that will not persist upon reboot.
	if err := sysctl.Set("kern.maxproc", "1000"); err != nil {
		t.Fatalf("error incrementing maxproc: %v\n", err)
	}
}

```
