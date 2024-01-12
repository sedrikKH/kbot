# kbot

## Задача 8.1

Реалізований pre-commit hook скрипт з автоматичним встановленням gitleaks залежно від операційної системи, з опцією enable за допомогою git config та інсталяцією за методом “curl pipe sh” (задача делегована junior та middle інженерам )

**Метод інсталяції був реалізовний junior та middle інженерами :-)**

Для ОС Linux - файл /script/gitleaks/install.sh 
 
```
curl -s https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install.sh | bash
```
 


Для ОС Windows - файл /script/gitleaks/install_gitleaks.ps1

```
powershell -c "Invoke-WebRequest -Uri https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install_gitleaks.ps1 -OutFile gitleaks.ps1; .\\gitleaks.ps1"
```

Для Darwin - файл /script/gitleaks/install.sh 
 
```
 curl -s https://raw.githubusercontent.com/sedrikKH/kbot/main/script/gitleaks/install.sh | bash

```

## Активація pre-commtit hook

Скопіювати файл pre-commit з папки script проекту у папку ./.git/hooks та зробити його виконуваним

```
cp ./script/pre-commit ./.git/hooks/
chmod 777 ./.git/hooks/pre-commit
```

## Тестовий приклад

```
git add .
git commit -m "Test commit"
```

Якщо в системі вже встановлений Gitleaks вивід буде наступним

```
git add .
git commit -m "Test commit"


    ○
    │╲
    │ ○
    ○ ░
    ░    gitleaks

9:42PM INF 1 commits scanned.
9:42PM INF scan completed in 8.48ms
9:42PM INF no leaks found
No secrets found, automatically accepting commit
Everything up-to-date
[main cc49737] Test commit
 1 file changed, 1 insertion(+)
```

**Якщо додати файл з токеном у каталог проекту та зробити комміт Gitleaks видаст помилку**

```
git commit -m "Test commit1"

    ○
    │╲
    │ ○
    ○ ░
    ░    gitleaks

Finding:     ...-from-literal=token="REDACTED"
Secret:      REDACTED
RuleID:      telegram-bot-api-token
Entropy:     4.466406
File:        secrets.yaml
Line:        1
Commit:      43e2b0c6cb4aa121f5b74f8267034cde67ef5a02
Author:      Sergiy
Email:       ada*********gmail.com
Date:        2024-01-12T21:48:40Z
Fingerprint: 43e2b0c6cb4aa121f5b74f8267034cde67ef5a02:secrets.yaml:telegram-bot-api-token:1

9:49PM INF 1 commits scanned.
9:49PM INF scan completed in 7.75ms
9:49PM WRN leaks found: 1
Secrets found, reverting commit
```

## Task 8.1 Done!









<!-- ![alt text](/img/kbot%20workflow-Page-2.drawio.png) -->

<!-- ## TELE_TOKEN

``` 
    read -s TELE_TOKEN 
    echo $TELE_TOKEN
    export TELE_TOKEN
```
## Add tags

```
git tag -a {тег} -m {комент}
```

## Build

Example:
``` 
    go build -ldflags "-X="hgithub.com/sedrikKH/prometheus_kbot/cmd.appVersion=v1.0.2 
```


## Start

```
./prometheus_kbot start

```
 -->

