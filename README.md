# kbot

![alt text](/img/kbot%20workflow-Page-2.drawio.png)

## TELE_TOKEN

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












