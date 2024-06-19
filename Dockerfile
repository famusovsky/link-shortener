FROM golang:1.22.2-alpine

WORKDIR /

ARG override_tables=false
ARG use_env=false
ENV OVERRIDE=$override_tables
ENV USE_ENV=$use_env

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o out ./cmd/web

EXPOSE ${port}

CMD [ "sh", "-c", "./out -override_tables=$OVERRIDE -use_env=$USE_ENV" ]