FROM golang:1.21.3

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY server/server.go ./server/
COPY html/head.html ./html/
COPY html/footer.html ./html/

RUN go build -o bin .

EXPOSE 9000

ENTRYPOINT [ "/app/bin" ]
