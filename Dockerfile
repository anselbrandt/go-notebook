FROM node:16 AS dependencies
WORKDIR /frontend
COPY ./frontend/package*.json ./
RUN npm install

FROM dependencies AS builder
COPY ./frontend .
RUN npm run build

FROM golang:alpine as gobuild
RUN apk add git
RUN apk add build-base
WORKDIR /src
COPY go.* ./
RUN go mod download
COPY *.go ./
COPY ./data ./data
COPY ./env ./env
COPY ./handlers ./handlers
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' .

FROM alpine:latest
COPY --from=gobuild ./src ./
COPY --from=builder ./frontend/build ./frontend/build
WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/go-notes"]