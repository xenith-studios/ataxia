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

all: install

%: install %.go
	$(QUOTED_GOBIN)/$(GC) $*.go
	$(QUOTED_GOBIN)/$(LD) -o $@ $*.$O

telnet.o: telnet.c
	gcc $(CGO_CFLAGS) $(_CGO_CFLAGS_$(GOARCH)) -fPIC $(CFLAGS) -c telnet.c -o telnet.o

examples: install
	cd example && make
