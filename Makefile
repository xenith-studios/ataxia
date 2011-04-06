# Explicitly make targets phony, just in case
.PHONY : all relconst deppkg pkgs cmd install clean nuke

# By default, build everything
all : deppkg pkgs cmd

relconst :
	@echo "Updating release constants..."
	tools/release-edit.sh

deppkg :
	@echo "Installing dependencies..."
	goinstall -u log4go.googlecode.com/hg

pkgs :
	@echo "Building packages..."
	$(MAKE) -C lib/telnet install
	$(MAKE) -C src/pkg/lua install
	$(MAKE) -C src/pkg/settings install
	$(MAKE) -C src/pkg/handler install

cmd : relconst
	@echo "Building ataxia..."
	$(MAKE) -C src/cmd/ataxia
	cp src/cmd/ataxia/ataxia bin/

helpers :
	@echo "Building helpers..."
	$(MAKE) -C src/cmd/gencert
	$(MAKE) -C src/cmd/stresstest

install : all
	$(MAKE) -C src/cmd/ataxia install

clean :
	$(MAKE) -C lib/telnet clean
	$(MAKE) -C src/pkg/lua clean
	$(MAKE) -C src/pkg/settings clean
	$(MAKE) -C src/pkg/handler clean
	$(MAKE) -C src/cmd/ataxia clean

nuke :
	$(MAKE) -C lib/telnet nuke
	$(MAKE) -C src/pkg/lua nuke
	$(MAKE) -C src/pkg/settings nuke
	$(MAKE) -C src/pkg/handler nuke
	$(MAKE) -C src/cmd/ataxia nuke
	rm -f bin/ataxia
