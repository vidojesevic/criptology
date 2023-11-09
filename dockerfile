FROM golang:1.21.3

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY server/server.go ./server/
COPY parser/parser.go ./parser/
COPY server/log/ ./server/log/
COPY web/views/head.html ./web/views/
COPY web/views/footer.html ./web/views/

RUN go build -o bin .

EXPOSE 9000

ENTRYPOINT [ "/app/bin" ]
