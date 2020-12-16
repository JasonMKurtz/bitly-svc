FROM golang:1 as builder
WORKDIR /build
COPY ./src/ /build
RUN CGO_ENABLED=0 go build -o /build/metrics endpoint.go 
RUN chmod +x /build/metrics

FROM golang:alpine 
WORKDIR /run
COPY --from=builder /build/metrics .
EXPOSE 5000
ENTRYPOINT ["/run/endpoint"]