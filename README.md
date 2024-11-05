## About FreedChat
FreedChat is a web based chat app written in angular and Go lang. You can chat anyone here as long as you login/register first.

## Entity Relational Diagram 
<img src="design/ERD.png" width="800" />

## Documentation for API Spec using Swagger/Open API
To view the api spec you can just right click and click preview the apispec.yaml that was inside "/internal" directory
<img src="design/ApiSpec.png" width="800" />

## Setup project
To setup the project you need to install this:
1. Golang Migrate: https://github.com/golang-migrate/migrate
2. Docker: https://docs.docker.com/engine/install/
3. Postgresql image:
```
docker pull postgres
```
4. Makefile

After that:
1. Create or make .env file inside "/internal" directory:
```
touch .env
```
2. Assign .env file like this for example
```
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_URL=postgresql://your_user:your_password@localhost:5432/freedchat?sslmode=disable
SECRET_KEY=your_secret_key
IMAGE_ENV=localhost:8081
```
3. Create docker compose
```
cd internal
docker compose create
```
4. Start the docker compose
```
docker compose start
```

## Run the project
To run the angular or front end (you have to be in root dir):
```
ng serve
```
To run go or backend:
```
cd internal
cd cmd
go run main.go
```
