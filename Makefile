# Builds the project
BINARY=csv2xml
# These are the values we want to pass for VERSION and BUILD
#git tag v2.6.0-beta
#git commit -am "One more change after the tags"
VERSION=2.6.0-beta
#VERSION=0.1beta
#BUILD=`date +%FT%T%z`
BUILD=`date +%F`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS_f1=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"
#LDFLAGS_f2=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Entry=f2"


build:
	GOARCH=amd64 GOOS=darwin go build ${LDFLAGS_f1} -o builds/${BINARY}-macosx
	GOARCH=arm64 GOOS=darwin go build ${LDFLAGS_f1} -o builds/${BINARY}-aarm64
	GOARCH=amd64 GOOS=linux go build  ${LDFLAGS_f1} -o builds/${BINARY}-linux
	GOARCH=amd64 GOOS=windows go build ${LDFLAGS_f1} -o builds/${BINARY}.exe
	GOARCH=386 GOOS=windows go build ${LDFLAGS_f1} -o builds/${BINARY}_x32.exe

run:
	./builds/${BINARY}-apple -2 file.csv

build_and_run: build run

app: build
	cp builds/${BINARY}.exe AutoFSA/
	cp builds/${BINARY}_x32.exe AutoFSA/

clean:
	go clean
	rm builds/${BINARY}-macosx
	rm builds/${BINARY}-aarm64
	rm builds/${BINARY}-linux
	rm builds/${BINARY}.exe
	rm builds/${BINARY}_x32.exe
	rm AutoFSA/${BINARY}_x32.exe
	rm AutoFSA/${BINARY}_x32.exe



