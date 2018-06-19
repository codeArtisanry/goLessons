// +build windows !linux !darwin !freebsd

package main

//go:generate goversioninfo -icon=icon.ico -o deploy.syso

// go build -ldflags -H=windowsgui
// CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w -linkmode internal"
fun hide(){

}