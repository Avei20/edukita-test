# Edukita Take Home Test

## Introduction
This is fullstack project using Go and Next.JS framework. The backend is written on minimalize Go backend. Frontend is used by Next.JS framework using shadcn/ui library.

## Library Used
- [UUID](https://pkg.go.dev/github.com/google/uuid)
- [JWT Library](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)
- [Next.JS](https://nextjs.org/)
- [Shadcn/ui](https://ui.shadcn.com/)

## Requirements
- [Docker](https://docs.docker.com/get-docker/)

## Breakdown of Documentation
- [Backend](backend/README.md)
- [Frontend](frontend/README.md)

## Development
1. Copy the env.template to .env
2. Run all the services using this command
```bash
docker compose up --build
```
4. If you want to hot-reload when developing all the services. Run this command
```bash
docker compose up --build --watch
```

## Deployment
If you want to deplay using docker-compose. Use the command below:
```bash
docker compose up --build
```
This command will be build the docker image and run the container.

For kubernetes or minikube deployment. Please refer to here [deployment](deployment/README.md)
