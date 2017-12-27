CARGO = cargo

CARGO_OPTS =

all: build-quick

full: build-full doc

build-quick: build-proxy build-engine

build-proxy:
	sh tools/release-edit.sh
	go build github.com/xenith-studios/ataxia/cmd/ataxia-proxy
	mv ataxia-proxy bin/

build-engine:
	$(CARGO) $(CARGO_OPTS) build --bin ataxia-engine
	cp -f target/debug/ataxia-engine bin/

build-full: full-build-proxy build-proxy full-build-engine build-engine

full-build-proxy:
	#dep ensure
	go fmt ./...
	goimports -w .
	go vet ./...
	golint ./... #| egrep -v "_string.go"

full-build-engine:
	$(CARGO) +nightly-2017-12-20 fmt
	env CARGO_TARGET_DIR=./target/clippy $(CARGO) +nightly clippy

clean:
	$(CARGO) $(CARGO_OPTS) clean
	rm -f bin/ataxia-{engine,proxy}

check:
	$(CARGO) $(CARGO_OPTS) check

test:
	$(CARGO) $(CARGO_OPTS) test

bench:
	$(CARGO) $(CARGO_OPTS) bench

doc:
	$(CARGO) $(CARGO_OPTS) doc

.PHONY: all build build-proxy build-engine clean check test bench doc
