FROM golang:1.20.12
LABEL authors="vladimir"

WORKDIR /Task1

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]