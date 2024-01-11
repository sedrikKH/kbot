
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# Устанавливаем git
choco install git -y
refreshenv

# Клонируем репозиторий
git clone https://github.com/gitleaks/gitleaks.git

# Собираем
Set-Location -Path .\gitleaks\
go generate ./...
go build -o gitleaks.exe

# Копируем бинарник в Program Files
$installPath = "$env:ProgramFiles\gitleaks" 
New-Item -ItemType Directory -Force -Path $installPath
Copy-Item -Path .\gitleaks.exe -Destination $installPath	

# Добавляем в PATH
$env:Path += ";$installPath"

# Проверка
gitleaks -v