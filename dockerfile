FROM golang:1.21.3

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY server/server.go ./server/
COPY datautil/* ./datautil/
COPY config/config.json ./config/
COPY logger/ ./logger/
COPY web/views/* ./web/views/
COPY web/scripts/* ./web/scripts/
COPY web/index.html ./web/

RUN go build -o bin .

EXPOSE 9000

ENTRYPOINT [ "/app/bin" ]
