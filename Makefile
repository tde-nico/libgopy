#GOFLAGS	= -a -installsuffix cgo -ldflags '-s'

all:
	@ make -s -C ./lib

run: all
	@ go test

test: run
