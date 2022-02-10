FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o webapp-hello .
CMD ["/app/webapp-hello"]
