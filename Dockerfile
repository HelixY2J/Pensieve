
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api *.go


FROM scratch
WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]