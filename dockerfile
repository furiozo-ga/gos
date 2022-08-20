FROM golang:1.19-alpine as build

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY gos.go ./
ENV CGO_ENABLED=0
RUN go build -o gos gos.go

FROM scratch
WORKDIR /app
COPY --from=build /build/gos .
EXPOSE 1111

CMD ["./gos"]