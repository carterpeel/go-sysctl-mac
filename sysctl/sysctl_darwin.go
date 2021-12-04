//+build darwin

package sysctl

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var (
	ErrorUnderPrivileged = fmt.Errorf("root/sudo is required for sysctl modification")
)

func Set(key string, value string) error {
	if os.Geteuid() != 0 {
		return ErrorUnderPrivileged
	}
	if out, err := exec.Command("sysctl", fmt.Sprintf("%s=%s", key, value)).CombinedOutput(); err != nil {
		return fmt.Errorf("error setting key:value pair: %v: %v", err, string(out))
	}
	return nil
}

func SetPersistent(key string, value string) error {
	if err := Set(key, value); err != nil {
		return err
	}
	fi, err := os.OpenFile("/private/etc/sysctl.conf", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func(fi *os.File) {
		_ = fi.Close()
	}(fi)
	if err := fi.Chown(0, 0); err != nil {
		return err
	}

	if err := setConfKey(key, value, fi); err != nil {
		return err
	}
	return nil
}

func setConfKey(key, value string, fi *os.File) error {
	var found bool
	var offset int
	for scanner := bufio.NewScanner(fi); scanner.Scan(); {
		if index := bytes.Index(scanner.Bytes(), []byte(key)); index > -1 {
			if _, err := fi.Seek(int64(offset), 0); err != nil {
				return err
			}
			if _, err := fi.WriteString(fmt.Sprintf("%s=%s", key, value)); err != nil {
				return err
			}
			found = true
			break
		}
		offset += len(scanner.Bytes())
	}
	if !found {
		if _, err := fi.Seek(0, 2); err != nil {
			return err
		}
		if _, err := fi.WriteString(fmt.Sprintf("%s=%s\n", key, value)); err != nil {
			return err
		}
	}
	return nil
}