package main

import (
	"fmt"

	"github.com/imroc/req"
)

var Execlogplan = "http://15.14.12.153:10040/logplan/api/v1/UtilityUtils/ACT_EXECPLANJOB?autoLoad=true&dbs=all&pwd=R3aB6x8f0"

func main() {

	// r := req.New()
	// r.SetFlags(req.LstdFlags | req.Lcost)
	resp2, err := req.Get(Execlogplan)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp2)

}
