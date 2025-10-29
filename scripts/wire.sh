#!/bin/bash
# Install wire tool
# check if wire is installed
if ! command -v wire &> /dev/null
then
    go install github.com/google/wire/cmd/wire@latest
fi
# Move to internal/app and run wire
cd internal/app
$(go env GOPATH)/bin/wire
