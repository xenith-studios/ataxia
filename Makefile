CARGO = cargo
CARGO_OPTS =

debug:
	$(CARGO) $(CARGO_OPTS) build
	cp -f target/debug/portal bin/portal
	cp -f target/debug/engine bin/engine

release:
	$(CARGO) $(CARGO_OPTS) build --release
	cp -f target/release/portal bin/portal
	cp -f target/release/engine bin/engine

full: lint debug doc

portal: lint
	$(CARGO) $(CARGO_OPTS) build --bin portal
	cp -f target/debug/portal bin/portal

engine: lint
	$(CARGO) $(CARGO_OPTS) build --bin engine
	cp -f target/debug/engine bin/engine

lint:
	$(CARGO) $(CARGO_OPTS) fmt
	env CARGO_TARGET_DIR=./target/clippy $(CARGO) $(CARGO_OPTS) clippy --workspace --all-targets

bootstrap:
	rustup component add --toolchain nightly rustfmt-preview
	rustup component add --toolchain nightly clippy-preview

clean:
	$(CARGO) $(CARGO_OPTS) clean
	rm -f bin/{engine,portal}

check:
	$(CARGO) $(CARGO_OPTS) check

test:
	$(CARGO) $(CARGO_OPTS) test

bench:
	$(CARGO) $(CARGO_OPTS) bench

doc:
	$(CARGO) $(CARGO_OPTS) doc

.PHONY: quick full portal engine lint bootstrap clean check test bench doc
