FROM golang:1.18

ENV GO111MODULE=on\
    GGOS=linux\
    GOARCH="amd64"\
    GOProxy="https://goproxy.cn,direct"\
    CGO_ENABLED=0

WORKDIR /project/go-docker/

# COPY go.mod,go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /project/go-docker/build/myapp .

EXPOSE 8080
ENTRYPOINT ["/project/go-docker/build/myapp"]