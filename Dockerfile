FROM golang:latest
COPY main.go .
# compile source code
RUN go build main.go

EXPOSE 8080

CMD ["./main"]