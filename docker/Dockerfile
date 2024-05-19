############################
# STEP 1 build executable binary
############################
FROM golang:1.22-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "${USER}"

WORKDIR $GOPATH/src/app
COPY go.mod go.sum ./

# Fetch dependencies.
RUN go mod download
RUN go mod verify

# Copy the source code.
COPY . .

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

############################
# STEP 2 build image
############################
FROM scratch

# Copy timezone (required for time.Now() to work correctly).
# By default, scatch images do not contain the timezone files.
COPY --from=builder /usr/share/zoneinfo/ /usr/share/zoneinfo/

# Import the certificate files for HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/app /go/bin/app

# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["/go/bin/app"]
