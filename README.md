# API RESTful Golang POC

Esta é uma prova de conceito de uma API RESTful em Go usando Gin, JWT, PostgreSQL, GORM, bcrypt e verificação de e-mail. Não inclui autenticação via Google.

## Endpoints
- POST /register
- POST /login
- GET /verify-email/:token
- GET /profile
- GET /users/:id
- GET /admin/dashboard

## Perfis de acesso
- admin: acesso total
- user: acesso próprio
- client: acesso ao próprio perfil e dados públicos

## Como rodar
1. Instale Go e PostgreSQL
2. Configure o banco de dados
3. Instale dependências: `go mod tidy`
4. Execute: `go run main.go`

## Estrutura
- main.go
- /controllers
- /middlewares
- /models
- /routes
- /utils

## Observações
- Use variáveis de ambiente para segredos
- SMTP pode ser simulado para POC


