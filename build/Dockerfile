FROM golang:1.22-alpine AS builder

# Defina o diretório de trabalho no container
WORKDIR /app

# Copie o módulo go e baixe as dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copie o código fonte do projeto
COPY . .

# Compile o aplicativo Go para um binário estático
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

# Estágio de execução
FROM alpine:3.19  

# Defina o diretório de trabalho
WORKDIR /root/

# Copie o binário compilado do estágio de construção
COPY --from=builder /app/cmd .

# Exponha a porta que sua aplicação utiliza
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./main"]

