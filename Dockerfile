FROM golang:alpine AS build
COPY src /src
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -o /bin/server /src/*.go
FROM scratch
COPY --from=build /bin/server /server
CMD ["/server"]
