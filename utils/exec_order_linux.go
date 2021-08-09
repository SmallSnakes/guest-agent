// +build !windows


package utils

import (
	"os/exec"
)

func ExecOrder(action string) (code int) {
	if action == "poweroff" || action == "reboot" {
		err := exec.Command("sh", "-c", action).Run()
		if err == nil {
			return 200
		} else {
			return 400
		}
	} else {
		return 403
	}
	//err := exec.Command("cmd", "/c", order).Run()
}
