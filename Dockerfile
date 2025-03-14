FROM ubuntu:latest
RUN apt update && \
    apt install -y golang

COPY . /app
WORKDIR /app

CMD ["go", "run", "main.go"]