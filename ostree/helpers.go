package ostree

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetBootedCommit() (commit string, err error) {
	var cmdOut []byte
	if cmdOut, err = exec.Command("ostree", "admin", "status").Output(); err != nil {
		return
	}

	outLines := strings.Split(string(cmdOut), "\n")
	for _, line := range outLines {
		if strings.HasPrefix(line, "*") {
			booted := strings.Split(line, " ")
			if len(booted) != 3 {
				err = fmt.Errorf("Unable to parse booted commit: %q", line)
				return
			}
			commit = strings.Split(booted[2], ".")[0]
		}
	}
	return
}

func GetTag(ref, tag string) (value string, err error) {
	var cmdOut []byte
	if cmdOut, err = exec.Command("ostree", "show", fmt.Sprintf("--print-metadata-key=%s", tag), ref).Output(); err != nil {
		return
	}
	value = strings.TrimSpace(string(cmdOut))
	value = strings.Trim(value, "'")
	return
}

func GetVersionTag(ref string) (version string, err error) {
	version, err = GetTag(ref, "version")
	return
}
