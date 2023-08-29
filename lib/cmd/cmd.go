package cmd

import (
	"os"
	"os/exec"
	"zrWorker/core/slog"
	"zrWorker/pkg/utils"
)

func GetVersion() string {

	pid := os.Getpid()

	cmd := exec.Command("ps", "-p", utils.GetInterfaceToString(pid), "-o", "comm=")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	return string(out)

}

func CleanLog() {

	cmd := exec.Command("bash", "-c", "find /zrtx/log/cyberspace  -mtime +1 -name \"*\" | xargs -I {} rm -rf {}")
	cmd1 := exec.Command("bash", "-c", "find /scanning-client  -mtime +1 -name \"*\" |grep worker| xargs -I {} rm -rf {}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	out1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		slog.Println(slog.DEBUG, err1)
	}

	slog.Println(slog.DEBUG, string(out), string(out1))
}
