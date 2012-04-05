# Explicitly make targets phony, just in case
.PHONY : all relconst deppkg pkgs cmd install clean

# By default, build everything
all : deppkg cmd

relconst :
	@echo "Updating release constants..."
	tools/release-edit.sh

deppkg :
	@echo "Installing dependencies..."
#	goinstall -u log4go.googlecode.com/hg

pkgs :
	@echo "Building packages..."
	$(MAKE) -C lib/telnet install
	$(MAKE) -C src/pkg/lua install
	$(MAKE) -C src/pkg/settings install
	$(MAKE) -C src/pkg/handler install

cmd : relconst
	@echo "Building ataxia..."
	go build
	mv ataxia bin/

helpers :
	@echo "Building helpers..."
	$(MAKE) -C src/cmd/gencert
	$(MAKE) -C src/cmd/stresstest

install : relconst
	go install

clean :
	rm -f bin/ataxia
