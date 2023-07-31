# Запуск

## Для запуска сервера на устройтсе должен быть установлен [Docker](https://docs.docker.com/get-docker/), [Docker-compose](https://docs.docker.com/compose/) и утилита [Make](https://www.gnu.org/software/make/manual/make.html).

## Порядок запуска : 

### 1. Перейти в корневую папку сервера cd softvare-engineering/server
### 2. Создать файл переменных окружения .env со следующими параметрами

PORT=:4000

PG_DB_HOST=postgres

PG_DB_PORT=5432

PG_DB_USER=postgres

PG_DB_PASS=postgres

PG_DB_NAME=web

MG_DB_USER=mongodb

MG_DB_HOST=mongodb  

MG_DB_PORT=27017

### 3. Написать команду make dc-build которая после установки запустит сервер на указанном в .env порту
### 4. После запуска сервера становится доступна интерактивная документация по ссылке http://localhost:4000/swagger/index.html#/



### Swagger api documentation

![image](https://user-images.githubusercontent.com/46971653/224574790-97b9675a-f08a-4163-9e14-cb39a0188149.png)


