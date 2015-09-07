all:
	go get github.com/jteeuwen/go-bindata/...
	go-bindata conf/
	go build
