FROM golang:1.16-alpine

WORKDIR /home/devdocker

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /docker-gs-ping
RUN go get github.com/gin-gonic/gin
RUN go install github.com/gin-gonic/gin

EXPOSE 9092

CMD [ "/docker-gs-ping" ]