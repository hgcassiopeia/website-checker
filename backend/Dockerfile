FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

ENV GOARCH=amd64

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /app
EXPOSE 3000
USER nonroot:nonroot
CMD ["/app"]