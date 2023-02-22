FROM golang:1.18.3-bullseye

WORKDIR /synapsis

COPY . .

RUN go mod download

RUN go build -o main

EXPOSE 3000

CMD
CMD ["./main"]