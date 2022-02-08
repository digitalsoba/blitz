FROM golang:1.17.6-alpine as webhook-builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod ./
COPY go.sum ./
COPY app.go ./
RUN go build github.com/adnanh/webhook && go build app

FROM alpine:3.10
WORKDIR /app
COPY --from=webhook-builder /app/webhook ./
COPY --from=webhook-builder /app/app ./
COPY ./hooks/hooks.yaml /app/hooks/hooks.yaml
CMD [ "/app/webhook", "-hooks", "/app/hooks/hooks.yaml", "-verbose" ]