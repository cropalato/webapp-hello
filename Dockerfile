FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main webapp-hello 
CMD ["/app/webapp-hello"]
