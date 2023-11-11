FROM golang:1.21.3

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY server/server.go ./server/
COPY datautil/* ./datautil/
COPY config/config.json ./config/
COPY logger/ ./logger/
COPY web/views/head.html ./web/views/
COPY web/views/footer.html ./web/views/

RUN go build -o bin .

EXPOSE 9000

ENTRYPOINT [ "/app/bin" ]
