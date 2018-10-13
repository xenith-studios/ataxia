CARGO = cargo
CARGO_OPTS = +nightly

build:
	$(CARGO) $(CARGO_OPTS) build
	cp -f target/debug/proxy bin/ataxia-proxy
	cp -f target/debug/engine bin/ataxia-engine

full: lint build doc

proxy: lint build-proxy
	$(CARGO) $(CARGO_OPTS) build --bin proxy
	cp -f target/debug/proxy bin/ataxia-proxy

engine: lint build-engine
	$(CARGO) $(CARGO_OPTS) build --bin engine
	cp -f target/debug/engine bin/ataxia-engine

lint:
	$(CARGO) $(CARGO_OPTS) fmt
	env CARGO_TARGET_DIR=./target/clippy $(CARGO) $(CARGO_OPTS) clippy

bootstrap:
	rustup component add --toolchain nightly rustfmt-preview
	rustup component add --toolchain nightly clippy-preview

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

.PHONY: quick full proxy engine build-proxy build-engine lint-proxy lint-engine bootstrap clean check test bench doc
