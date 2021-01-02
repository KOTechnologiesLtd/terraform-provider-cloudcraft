TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
HOSTNAME=github.com
NAMESPACE=KOTechnologiesLtd
NAME=cloudcraft
BINARY=terraform-provider-${NAME}
PLUGINDIR=~/.terraform.d/plugins
VERSION=1.0.0
OS_ARCH=linux_amd64

default: build

init:
	go mod init ${BINARY}
	#GOPRIVATE=github.com/KOTechnologiesLtd/go-cloudcraft-api go mod vendor
	go mod vendor

reinit:
	rm -f go.sum go.mod
	rm -rf vendor
	go mod init ${BINARY}
	##GOPRIVATE=github.com/KOTechnologiesLtd/go-cloudcraft-api go mod vendor
	go mod vendor

build: fmtcheck
	go build ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install:
	#<TF13
	mkdir -vp $(PLUGINDIR)
	cp -f ${BINARY} $(PLUGINDIR)/${BINARY}
	#>TF13
	mkdir -p ${PLUGINDIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ${PLUGINDIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}_${VERSION}_${OS_ARCH}

rebuildinstall: build install

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
	CLOUDCRAFT_APITOKEN=fake RECORD=false TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout=10m

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

documentation:
	./tfplugindocs generate