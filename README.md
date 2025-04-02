# golang-banking
A simple banking service written in GoLang using Postgres, Redis, Gin, and gRPC, AWS and deployed on kubernetes.



docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine