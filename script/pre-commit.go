package main

import (  
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

func main() {

    installGitleaks() 
    
    // дальнейшая работа с gitleaks   
}




