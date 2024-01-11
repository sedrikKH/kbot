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

else

  echo "Unsupported OS"
  exit 1

fi

pip install pre-commit

# Далее установка...

if [[ "$OSTYPE" == "linux-gnu"* ]]; then

  # Linux
  git clone https://github.com/gitleaks/gitleaks.git
  cd gitleaks
  make build
  sudo cp gitleaks /usr/local/bin  
  sudo chmod +x /usr/local/bin/gitleaks
  gitleaks version

elif [[ "$OSTYPE" == "darwin"* ]]; then

  # MacOS  
  git clone https://github.com/gitleaks/gitleaks.git
  cd gitleaks
  make build
  gitleaks version
  
else
  echo "Unknown OS"
  exit 1  
fi

# Cleanup
cd .. 
rm -rf gitleaks

echo "git config --global --add gitleaks.enabled true"
git config --global --add gitleaks.enabled true

#$HOME/.local/bin/pre-commit install

