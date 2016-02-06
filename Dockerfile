# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
RUN go get github.com/tools/godep
RUN CGO_ENABLED=0 go install -a std

# Copy the local package files to the container's workspace.
ADD . /root/workarea/src/person

RUN godep go install root/workarea/src/person

# Run the outyet command by default when the container starts.
# Need to changes it to docker ENV
ENTRYPOINT root/workarea/bin/person -config=/root/workarea/src/person/config/auth-config-dev.ini

# Document that the service listens on port 9292.
EXPOSE 9292