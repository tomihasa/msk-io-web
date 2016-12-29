TOP := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT = $(realpath $(TOP))

GOPATH = $(ROOT)
export GOPATH

# Lint all code
lint:
	golint .

# Install all the dependencies that are shared across modules
deps:
	goapp get google.golang.org/appengine
	goapp get cloud.google.com/go/trace

# Serve via a local instance
serve:
	goapp serve app.yaml

deploy: 
	goapp deploy -application msk-io-web -version beta 