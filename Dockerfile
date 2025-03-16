FROM ubuntu:latest

RUN apt update && \
    apt install -y ca-certificates golang && \
    update-ca-certificates

WORKDIR /app

CMD ["go", "run", "cmd/server/main.go"]