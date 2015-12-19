PROJECT := gitlab2gogs
SCRIPTDIR := $(shell pwd)

GOBUILDDIR := $(SCRIPTDIR)/.gobuild
SRCDIR := $(SCRIPTDIR)
BINDIR := $(SCRIPTDIR)

ORGPATH := github.com/ewoutp
ORGDIR := $(GOBUILDDIR)/src/$(ORGPATH)
REPONAME := $(PROJECT)
REPODIR := $(ORGDIR)/$(REPONAME)
REPOPATH := $(ORGPATH)/$(REPONAME)
BIN := $(BINDIR)/$(PROJECT)

GOPATH := $(GOBUILDDIR)

SOURCES := $(shell find $(SRCDIR) -name '*.go')

ifndef GOOS
	GOOS := $(shell go env GOOS)
endif
ifndef GOARCH
	GOARCH := $(shell go env GOARCH)
endif


.PHONY: clean test

all: $(BIN)

clean:
	rm -Rf $(BIN) $(GOBUILDDIR)

.gobuild:
	mkdir -p $(ORGDIR)
	rm -f $(REPODIR) && ln -s ../../../../src $(REPODIR)
	git clone git@github.com:ewoutp/go-gitlab-client.git $(GOBUILDDIR)/src/github.com/ewoutp/go-gitlab-client
	git clone git@github.com:gogits/go-gogs-client.git $(GOBUILDDIR)/src/github.com/gogits/go-gogs-client

$(BIN): .gobuild $(SOURCES)
	go build -a -o $(PROJECT)
