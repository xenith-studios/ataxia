CARGO = cargo

CARGO_OPTS =

all: build doc

build: build-proxy build-engine

build-proxy:
	$(CARGO) fmt
	$(CARGO) $(CARGO_OPTS) build --bin ataxia-proxy
	cp -f target/debug/ataxia-proxy bin/

build-engine:
	$(CARGO) fmt
	$(CARGO) $(CARGO_OPTS) build --bin ataxia-engine
	cp -f target/debug/ataxia-engine bin/

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
