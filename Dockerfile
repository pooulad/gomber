FROM golang:latest

WORKDIR /gomber

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["make","run"]
