CARGO = cargo
GO = go

CARGO_OPTS =

all: build-quick

full: build-full doc

proxy: lint-proxy build-proxy

engine: lint-engine build-engine

build-quick: build-proxy build-engine

build-full: lint-proxy build-proxy lint-engine build-engine

build-proxy:
	sh tools/release-edit.sh
	$(GO) build github.com/xenith-studios/ataxia/cmd/proxy
	mv proxy bin/ataxia-proxy

build-engine:
	$(CARGO) $(CARGO_OPTS) build --bin engine
	cp -f target/debug/engine bin/ataxia-engine

lint-proxy:
	goimports -w .
	$(GO) vet ./...
	golint {cmd/proxy,engine,handler,utils}

lint-engine:
	$(CARGO) +nightly fmt
	env CARGO_TARGET_DIR=./target/clippy $(CARGO) +nightly clippy

bootstrap:
	dep ensure

clean:
	$(CARGO) $(CARGO_OPTS) clean
	rm -f bin/ataxia-{engine,proxy}

check:
	$(GO) build github.com/xenith-studios/ataxia/cmd/proxy
	rm -f proxy
	$(CARGO) $(CARGO_OPTS) check

test:
	$(GO) test github.com/xenith-studios/ataxia/{cmd/proxy,engine,handler,utils}
	$(CARGO) $(CARGO_OPTS) test

bench:
	$(CARGO) $(CARGO_OPTS) bench

doc:
	$(CARGO) $(CARGO_OPTS) doc

.PHONY: all build build-proxy build-engine clean check test bench doc
