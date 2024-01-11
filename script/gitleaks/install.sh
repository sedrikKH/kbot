#!/bin/bash

# Проверяем ОС 
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
   
  # Ubuntu/Debian
  if [[ -f /etc/debian_version ]]; then
    sudo apt update 
    sudo apt install -y git make python3-pip

  fi
  
  # CentOS/RHEL
  if [[ -f /etc/redhat-release ]]; then
    sudo yum update 
    sudo yum install -y git make python3-pip
  fi

elif [[ "$OSTYPE" == "darwin"* ]]; then
  
  # Mac OS
  if [[ -f /System/Library/CoreServices/SystemVersion.plist ]]; then
    sudo brew update
    sudo brew install git make pip3
  fi

elif [[ "$OSTYPE" = "cygwin" ]] || [[ "$OSTYPE" = "msys" ]] || [[ "$OSTYPE" = "win32" ]]; then

  # Установка утилит на Windows
  echo "Installing utilities for Windows"

  choco install -y git make wget unzip python pip

  refreshenv

else

  echo "Unsupported OS"
  exit 1

fi

pip install pre-commit

$HOME/.local/bin/pre-commit install

# Далее установка...

git clone https://github.com/gitleaks/gitleaks.git
cd gitleaks
make build
sudo cp gitleaks /usr/local/bin  
sudo chmod +x /usr/local/bin/gitleaks
gitleaks --version

# Cleanup
cd .. 
rm -rf gitleaks