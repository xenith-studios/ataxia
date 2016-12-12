CARGO = cargo

CARGO_OPTS =

all:
	$(MAKE) build
	$(MAKE) doc

build:
	$(CARGO) fmt
	$(CARGO) $(CARGO_OPTS) build
	cp -f target/debug/ataxia-{engine,proxy} bin/

clean:
	$(CARGO) $(CARGO_OPTS) clean
	rm -f bin/ataxia-{engine,proxy}

check:
	$(MAKE) build
	$(MAKE) test

test:
	$(CARGO) $(CARGO_OPTS) test

bench:
	$(CARGO) $(CARGO_OPTS) bench

doc:
	$(CARGO) $(CARGO_OPTS) doc

.PHONY: all build clean check test bench doc
