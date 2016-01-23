.PHONY: build

VERSION := "0.0.1"
BUILDSTAMP :=`date -u '+%Y%m%d-%I%M%S'`
GITHASH := `git rev-parse HEAD`


GOOS:="linux"
# GOOS:="darwin"
GOARCH:="386"

defult: build

build: 
	go build -ldflags "-X main.CfgBuildStamp=$(BUILDSTAMP) -X main.CfgVersion=$(VERSION) -X main.CfgGitHash=$(GITHASH)" -v -o ./bin/app.bin ./

build_linux_386:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -ldflags "-X main.CfgBuildStamp=$(BUILDSTAMP) -X main.CfgVersion=$(VERSION) -X main.CfgGitHash=$(GITHASH)" -v -o ./bin/app.bin ./