package IsWadmin

import (
	"os/exec"
	"runtime"
)

func IsWadmin() bool {
	if runtime.GOOS != "windows" {
		return false
	}

	if err := exec.Command("fsutil", "dirty", "query", "%systemdrive%").Run(); err != nil {
		if err := exec.Command("fltmc").Run(); err != nil {
			return false
		}
	}

	return true
}
