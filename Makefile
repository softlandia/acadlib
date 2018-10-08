PACKAGE := $(shell go list)
GOOS := $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
OBJ_DIR := $(GOPATH)/pkg/$(GOOS)_$(GOARCH)/$(PACKAGE)
