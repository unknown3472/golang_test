package git

import (
	"os/exec"
)
func GetUserName() (string, error) {
	cmd := exec.Command("git", "config", "--global", "user.name")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
func GetUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "--global", "user.email")
	output, err := cmd.Output()
	if err != nil {
		return "", err	
	}
	return string(output), nil
}
