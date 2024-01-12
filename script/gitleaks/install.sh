#!/bin/bash

# Проверяем ОС 
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
   
  # Ubuntu/Debian
  if [[ -f /etc/debian_version ]]; then
    sudo apt update 
    sudo apt install -y git make 
  fi
  
  # CentOS/RHEL
  if [[ -f /etc/redhat-release ]]; then
    sudo yum update 
    sudo yum install -y git make 
  fi

elif [[ "$OSTYPE" == "darwin"* ]]; then
  
  # Mac OS
  if [[ -f /System/Library/CoreServices/SystemVersion.plist ]]; then
    sudo brew update
    sudo brew install git make 
  fi

else

  echo "Unsupported OS"
  exit 1

fi

# Далее установка...

  git clone https://github.com/gitleaks/gitleaks.git
  cd gitleaks
  make build

if [[ "$OSTYPE" == "linux-gnu"* ]]; then

  # Linux
  sudo cp gitleaks /usr/local/bin  
  sudo chmod +x /usr/local/bin/gitleaks
  gitleaks version

elif [[ "$OSTYPE" == "darwin"* ]]; then

  # MacOS  
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



