// +build windows


package utils

import (
	"os/exec"

)

func ExecOrder(action string) (code int) {
	if action == "poweroff" {
		err := exec.Command("cmd", "/c", "shutdown -s -t 0").Run()

		if err == nil {
			return 200
		} else {
			return 400
		}
	}else if action == "reboot" {
		err := exec.Command("cmd", "/c", "shutdown -r -t 0").Run()

		if err == nil {
			return 200
		} else {
			return 400
		}
	}else {
		return 403
	}
}
