FROM golang:1.20.1-alpine

RUN mkdir /app
ADD . /app

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . /

RUN go build -o /Assignment3Go

EXPOSE 8000

CMD [ "/Assignment3Go" ]