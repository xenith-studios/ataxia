# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

CGO_CFLAGS+=-Ilibtelnet -DHAVE_ZLIB
CGO_DEPS=_cgo_export.o
CGO_LDFLAGS+=-lz
CGO_OFILES=\
		telnet.o \
		libtelnet/libtelnet.o

TARG=telnet

CGOFILES=\
		telnet.go \

CLEANFILES+=\
		libtelnet/libtelnet.o

include $(GOROOT)/src/Make.pkg

all: install examples

%: install %.go
	$(QUOTED_GOBIN)/$(GC) $*.go
	$(QUOTED_GOBIN)/$(LD) -o $@ $*.$O

golua.o: golua.c
	gcc $(CGO_CFLAGS) $(_CGO_CFLAGS_$(GOARCH)) -fPIC $(CFLAGS) -c golua.c -o golua.o

lua_defs.go:
	echo "package golua;" > lua_defs.go
	echo "$(LUA_INCLUDE_DIRECTIVES)" "import \"C\"" >> lua_defs.go
#	echo "import \"C\"" >> lua_defs.go
	cat $(LUA_HEADER_FILES) | grep '#define LUA' | sed 's/#define/const/' | sed 's/\([A-Z_][A-Z_]*\)[[:space:]]*.*/\1 = C.\1/'  >> lua_defs.go

examples: install
	cd example && make
