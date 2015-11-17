# Starting from the latest Golang container
FROM golang:1.5.1
MAINTAINER Pavel Pavlov

# Set the WORKDIR to the project path in your GOPATH, e.g. /go/src/github.com/go-martini/martini/
WORKDIR $GOPATH/src/app

# ADD the content of your repository into the container
COPY . ./

# Install dependencies through go get, unless you vendored them in your repository before
# Vendoring can be done through the godeps tool or Go vendoring available with
RUN go get github.com/tools/godep
RUN godep go build

EXPOSE 3002

CMD ["./app", "-port", "3002"]
