FROM golang:1.25
RUN go install github.com/air-verse/air@latest
COPY . /app
WORKDIR /app
EXPOSE 3001
ENTRYPOINT ["air"]