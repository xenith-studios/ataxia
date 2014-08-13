# Explicitly make targets phony, just in case
.PHONY : all relconst deppkg cmd install clean

# By default, build everything
all : deppkg cmd

relconst :
	@echo "Updating release constants..."
	tools/release-edit.sh

deppkg :
	@echo "Installing dependencies..."
	go get -u github.com/aarzilli/golua/lua
	go get -u github.com/stevedonovan/luar
	go get -u github.com/nu7hatch/gouuid
#	go get -u code.google.com/p/log4go

cmd : relconst
	@echo "Building ataxia..."
	go build
	mv ataxia bin/

helpers :
	@echo "Building helpers..."
	go build gencert
	go build stresstest

install : relconst
	go install

clean :
	rm -f bin/ataxia
