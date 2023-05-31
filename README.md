## Структура проекта
- **config.yml** - конфиг, который запускает все необходимые сущности
    - в папке "configs" располагаются отдельные конфиги
    - скрипт "generate" собирает конфиги в один (сам config.yml) и генерирует сущности для сервисов "cats" и "dogs"
- Dockerfile, который собирает stateful-приложение из api.go
    - [Ссылка на образ контейнера nikigor/sre:3](https://hub.docker.com/layers/nikigor/sre/3/images/sha256-69753660099a27546f27d8cdc29d36df7b2454e7b18c6a42c7185e7ab3572efc?tab=layers)

## Ingress работает с хостом ```app.example.com```
Чтобы поменять его у себя на машине, необходимо
1. ```
   kubectl apply -f config.yml 
   ```
1. ```
   kubectl describe Ingress
   ```
   нужно взять поле `Address`

1. Добавить в `/etc/hosts`
    ```
   192.168.121.2	app.example.com
    ```
    вместо `192.168.121.2` нужно подставить свой Ingress Address
