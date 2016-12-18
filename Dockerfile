# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/newlee/petstore
ADD ./build/build-in-docker.sh /go/build.sh

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
#RUN go test github.com/newlee/petstore/server && go build github.com/newlee/petstore

