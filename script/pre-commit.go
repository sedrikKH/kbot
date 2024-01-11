package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func installGitleaks() {

	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "linux":
		cmd = exec.Command("bash", "-c",
			"curl -s https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install.sh | bash")

	case "windows":
		cmd = exec.Command("powershell", "-c",
			"Invoke-WebRequest -Uri https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install_gitleaks.ps1 -OutFile gitleaks.ps1; .\\gitleaks.ps1")

	case "darwin":
		cmd = exec.Command("bash", "-c",
			"curl -s https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install.sh | bash")

	default:
		panic("Unsupported OS")
	}

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func checkLastCommit() {

	GITLEAKS_REPORT := "report.json"
	GITLEAKS_OPTS := "detect --redact -v"
	GITLEAKS_GIT_LOGS := "HEAD~1^..HEAD"

	out, err := exec.Command("gitleaks", GITLEAKS_OPTS, GITLEAKS_REPORT, GITLEAKS_GIT_LOGS).Output()

	if err != nil {
		//panic(
		fmt.Println(err)
	}

	if string(out) != "" {
		// найдены секреты - откатываем коммит
		exec.Command("git", "reset", "HEAD~1").Run()
	} else {
		// принимаем коммит, если чистый
		fmt.Println("No secrets found, automatically accepting commit.")
		exec.Command("git", "commit", "--amend", "--no-edit").Run()
		exec.Command("git", "push").Run()
	}
}

func main() {

	installGitleaks()

	checkLastCommit()

}
