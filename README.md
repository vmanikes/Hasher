# Response Hasher

A CLI tool that takes in http url's in the commandline and returns the hash of the response of the url's when a http request
is made to them

This tool can also make requests in parallel by taking in the `-parallel` flag

## How to run
```shell
go build main.go
./main -parallel 3 https://www.google.com https://www.facebook.com https://www.twitter.com
```

### Expected output
```shell
https://google.com adf73aa6ad31207f2f3712577004bcce
https://wwww.facebook.com 4c186b87c6e43566e795279aa6f8d9e1
https://www.twitter.com 6af1eda889ce514f31a2349b12e1907f
```

## Running tests
`go test -v ./...`
