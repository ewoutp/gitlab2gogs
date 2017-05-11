PROJECT := gitlab2gogs
SCRIPTDIR := $(shell pwd)

GOBUILDDIR := $(SCRIPTDIR)/.gobuild
SRCDIR := $(SCRIPTDIR)
BINDIR := $(SCRIPTDIR)

ORGPATH := github.com/ewoutp
ORGDIR := $(GOBUILDDIR)/src/$(ORGPATH)
REPONAME := $(PROJECT)
REPODIR := $(ORGDIR)/$(REPONAME)
BIN := $(BINDIR)/$(PROJECT)

GOPATH := $(GOBUILDDIR)

SOURCES := $(shell find $(SRCDIR) -name '*.go')

.PHONY: clean test

all: $(BIN)

clean:
	rm -Rf $(BIN) $(GOBUILDDIR)

.gobuild:
	mkdir -p $(ORGDIR)
	rm -f $(REPODIR) && ln -s ../../../../src $(REPODIR)
	git clone https://github.com/xanzy/go-gitlab.git $(GOBUILDDIR)/src/github.com/xanzy/go-gitlab -b f-api-v4
	git clone https://github.com/gogits/go-gogs-client.git $(GOBUILDDIR)/src/github.com/gogits/go-gogs-client
	git clone https://github.com/google/go-querystring.git $(GOBUILDDIR)/src/github.com/google/go-querystring

$(BIN): .gobuild $(SOURCES)
	go build -a -o $(PROJECT)
