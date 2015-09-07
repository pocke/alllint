all: depends
	go build

depends:
	go get github.com/jteeuwen/go-bindata/...
	go-bindata conf/
	go get

test: depends
	go test -v --race
