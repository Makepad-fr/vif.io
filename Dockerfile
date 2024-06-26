FROM golang:1.22 as go-builder

WORKDIR /app

COPY Makefile ./Makefile
COPY go.work go.work

COPY server/ ./server/
COPY services/ ./services/
COPY model/ ./model/

ENV CGO_ENABLED=0

RUN make build

FROM alpine:3.20

WORKDIR /app
EXPOSE 4321
COPY --from=go-builder /app/out/server /app/server
COPY templates/ ./templates
COPY static/ ./static
ENTRYPOINT [ "/app/server" ]