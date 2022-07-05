FROM golang:1.18.3

WORKDIR /app

COPY . .

RUN go get -u github.com/cosmtrek/air

CMD ["go", "run", "main.go"]
# CMD ["air", "-c", ".air.toml"]
