FROM golang:1.15.8-alpine3.12
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git
RUN go get -d github.com/bwmarrin/discordgo
RUN go get -d github.com/joho/godotenv
RUN go build -o main .
CMD ["/app/main"]