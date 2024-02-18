FROM golang:alpine3.19 AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build -o score main.go

FROM scratch

WORKDIR /build

COPY --from=builder /build/score /build/score

ENTRYPOINT ["./score"]