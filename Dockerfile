FROM golang:1.13 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY droxo ./droxo

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/blog .

EXPOSE 8102

ENTRYPOINT [ "./blog" ]
