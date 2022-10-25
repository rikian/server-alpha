FROM golang

WORKDIR /workdir/server-alpha

COPY . .

RUN go mod tidy
RUN go build -o ./server-alpha .

CMD ["./server-alpha"]