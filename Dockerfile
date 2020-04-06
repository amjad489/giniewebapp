FROM golang:1.13
WORKDIR /go/src/giniewebapp
COPY . .

RUN go get -d -v ./...
RUN go build
RUN ls ./
CMD ["./giniewebapp"]