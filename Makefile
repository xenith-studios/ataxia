CARGO = cargo
GO = go
CARGO_OPTS =

quick: build-proxy build-engine

full: proxy engine doc

proxy: lint-proxy build-proxy

engine: lint-engine build-engine

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
	golint {cmd/proxy,internal/*}

lint-engine:
	$(CARGO) $(CARGO_OPTS) fmt
	env CARGO_TARGET_DIR=./target/clippy $(CARGO) $(CARGO_OPTS) +nightly clippy

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
	$(GO) test ./...
	$(CARGO) $(CARGO_OPTS) test

bench:
	$(CARGO) $(CARGO_OPTS) bench

doc:
	$(CARGO) $(CARGO_OPTS) doc

.PHONY: quick full proxy engine build-proxy build-engine lint-proxy lint-engine bootstrap clean check test bench doc
