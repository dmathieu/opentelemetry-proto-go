#!/bin/bash -e

DIR=`pwd`
TMP_DIR=`mktemp -d`

git clone https://github.com/open-telemetry/opentelemetry-proto $TMP_DIR
cd $TMP_DIR

for file in $(find . -name '*.proto'); do
	sed -i 's,go_package = "github.com/open-telemetry/opentelemetry-proto/gen/go,go_package = "github.com/dmathieu/opentelemetry-proto-go,' $file
done

sed -i 's,gogo_,gogofaster_,' Makefile
make gen-go

cp -R $TMP_DIR/gen/go/github.com/dmathieu/opentelemetry-proto-go/* $DIR
