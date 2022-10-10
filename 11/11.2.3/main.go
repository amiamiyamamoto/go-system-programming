package main

import "os/exec"

var cmd *exec.Cmd

if runtime.GOOS == "windows" {
	cmd == exec.Command("cmd.exe", "/C", "timeout 5")
}
cmd.Run()