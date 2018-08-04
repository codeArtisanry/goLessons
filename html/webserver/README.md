# Webserver [![GoDoc](https://img.shields.io/badge/go-documentation-orange.svg)](https://godoc.org/github.com/tahasevim/webserver)[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://travis-ci.org/tahasevim/webserver)
Simple webserver that serves with localhost at `port 8080` as a default port.<br>
To initialize webserver with different port, use `port` flag.<br>
## Install
```bash
go get github.com/tahasevim/webserver
```
## Usage
To run webserver without `port` flag, type below command in terminal:
```bash
webserver
```
To run with a custom port use `port` flag:
```bash
webserver -port=YourPortNumber
```
`YourPortNumber` should be chosen carefully since first 1000 `port` may be reserved for specific applications.
