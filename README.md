# Get Up and Running for Development

1. Have golang installed.
2. Run:

```sh
sudo apt-get install python3.11 python3.11-venv
python3 auto.py setup
```

## About template

This project has a [gateway](./gateway/) + a microservice called [auth](./auth/).

This project can be used to get up and running as fast as possible with a microservice project.

# How to Run Microservices:

## 1. Gateway

```sh
cd gateway
go run main.go
```

## 2. Auth Microservice

```sh
cd auth
go run main.go
```
