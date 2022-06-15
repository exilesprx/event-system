FROM golang:1.18.3 as build

WORKDIR /go/src

COPY ./ /go/src/

RUN go get -d

RUN go build -o build/app

FROM golang:1.18.3 as app

WORKDIR /go/src

COPY --from=build /go/src/build/ /go/src/

CMD ./app