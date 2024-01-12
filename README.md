# kbot

## Задача 8.1

реалізований pre-commit hook скрипт з автоматичним встановленням gitleaks залежно від операційної системи, з опцією enable за допомогою git config та інсталяцією за методом “curl pipe sh” (задача делегована junior та middle інженерам )

Метод інсталяції був реалізовний інженерами :-)

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

