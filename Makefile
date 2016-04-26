
# makefile to build and test Luke's go program
# if you want to build for other targets
#
# make target-name 
#

ARCH=amd64
PACKAGE=luke

$(TARGET):
	GOOS=$(TARGET) GOARCH=$(ARCH) go install luke
	GOOS=$(TARGET) GOARCH=$(ARCH) go build -o main.$(TARGET) luke.go 

all:
	for targ in windows linux darwin; do \
		export TARGET=$$targ ; \
		make $$targ ; \
	done


test:
	@(cd src/luke; go test -v $(PACKAGE)) || echo "damn! testing $(PACKAGE) failed"

