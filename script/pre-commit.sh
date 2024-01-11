#!/bin/bash

# Установка gitleaks
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  curl -s https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install.sh | bash
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "cygwin" ]]; then
  powershell -c "Invoke-WebRequest -Uri https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install_gitleaks.ps1 -OutFile gitleaks.ps1; .\\gitleaks.ps1"  
elif [[ "$OSTYPE" == "darwin"* ]]; then
  curl -s https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install.sh | bash
else
  echo "Unknown OS"
  exit 1
fi

# Проверка последнего коммита
GITLEAKS_REPORT="report.json"
GITLEAKS_OPTS="--redact -v"
GITLEAKS_GIT_LOGS="HEAD~1..HEAD"

#gitleaks detect --redact -v --report-path report.json --log-opts="HEAD~1..HEAD"
gitleaks detect  $GITLEAKS_OPTS --report-path $GITLEAKS_REPORT --log-opts="$GITLEAKS_GIT_LOGS"


if [ $? -eq 0 ]; then
  echo "No secrets found, automatically accepting commit"
  git commit --amend --no-edit
  git push
else
  echo "Secrets found, reverting commit" 
  git reset HEAD~1
fi