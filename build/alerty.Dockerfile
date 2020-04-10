FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


# Final stage
FROM scratch

COPY build/alerty_entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY --from=builder /app/ /app/

ENTRYPOINT ["/entrypoint.sh"]
CMD [ "brain" ]
