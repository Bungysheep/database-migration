FROM golang

WORKDIR /app/database-migration

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./bin/database-migration ./cmd

CMD ["./bin/database-migration"]