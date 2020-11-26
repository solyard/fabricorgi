FROM golang:alpine as build-env
LABEL maintainer="dizstorm@gmail.com"
ENV CGO_ENABLED=0
COPY . /app
WORKDIR /app
RUN go mod download && go build -o /fabricorgi

FROM hyperledger/fabric-tools:2.0 as main
COPY --from=build-env /fabricorgi /usr/bin/fabricorgi
ENTRYPOINT [ "fabricorgi" ]
