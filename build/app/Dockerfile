FROM golang:1.16.5 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/ishihaya/company-official-app-backend

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build .

FROM gcr.io/distroless/static-debian10 AS deploy
COPY --from=build /go/src/github.com/ishihaya/company-official-app-backend /

EXPOSE 8080

ENTRYPOINT [ "/company-official-app-backend" ]