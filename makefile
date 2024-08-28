# basic info
app := lobby
module := github.com/dstgo/lobby/cmd/lobby
# meta info
author = 246859
build_time := $(shell date +"%Y.%m%d.%H%M%S")
git_version := $(shell git tag --sort=-version:refname | sed -n 1p)

# build info
mode := debug
output := $(shell pwd)/bin
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)
os := $(host_os)
arch := $(host_arch)
ldflags := $(nullstring)

# reduce binary size at release mode
ifeq ($(mode), release)
	ldflags += -s -w
endif

# inject meta info
ifneq ($(app), $(nullstring))
	ldflags += -X main.AppName=$(app)
endif
ifneq ($(author), $(nullstring))
	ldflags += -X main.Author=$(author)
endif
ifneq ($(build_time), $(nullstring))
	ldflags += -X main.BuildTime=$(build_time)
endif
ifneq ($(git_version), $(nullstring))
	ldflags += -X main.Version=$(git_version)
endif

# binary extension
exe = $(nullstring)
ifeq ($(os), windows)
	exe = .exe
endif

.PHONY: build
build:
	# go lint
	go vet ./...

	# prepare target environment $(os)/$(arch)
	go env -w GOOS=$(os)
	go env -w GOARCH=$(arch)

	# build go module
	go build -trimpath \
		-ldflags="$(ldflags)" \
		-o $(output)/$(mode)/$(app)-$(os)-$(arch)/$(app)$(exe) \
		$(module)

	# resume host environment $(host_os)/$(host_arch)
	go env -w GOOS=$(host_os)
	go env -w GOARCH=$(host_arch)


# support platforms
windows := 386 amd64 arm64 arm
linux := 386 amd64 arm64 arm
darwin := amd64 arm64
platforms := windows linux darwin

.PHONY: build_all
build_all:
	@$(foreach os_i, $(platforms), \
		$(foreach arch_j, $(call $(os_i)), \
			$(shell $(MAKE) build os=$(os_i) arch=$(arch_j) mode=$(mode))))

# ent schema path
schema =
ent_out := ./server/data/ent
ent_target := schema
ent_generated := $(shell find $(ent_out)/* ! -path "*$(ent_target)*")

.PHONY: ent_gen
ent_gen:
ifneq ($(schema),)
	# generate new $(schema) schema
	ent new --target $(ent_out)/$(ent_target) $(schema)
endif
	# generate ent code
	ent generate $(ent_out)/$(ent_target)

.PHONY: ent_clean
ent_clean:
	@rm -rf $(ent_generated)

api_path := ./server/api

.PHONY: swag
swag:
	go generate $(api_path)


wire_out := ./server/

.PHONY: wire
wire:
	# generate app dependencies injection file
	wire gen $(wire_out)