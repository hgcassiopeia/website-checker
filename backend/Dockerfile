FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server

CMD ["./server"]

# ENV GOARCH=amd64

# FROM gcr.io/distroless/base-debian11
# COPY --from=build /go/bin/app /app
# EXPOSE 3000
# USER nonroot:nonroot
# CMD ["/app"]