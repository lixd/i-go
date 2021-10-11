.PHONY: build tools

all:
	@echo "make <tools|build|etc...>"

print-%: ; @echo $*=$($*)

domains:
	curl https://publicsuffix.org/list/effective_tld_names.dat > tldomains.dat

##
## Building
##
dist: domains
	@go generate
	@go build -i

build:
	go generate
	go build -i

test:
	go test
