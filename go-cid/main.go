package main

import (
	"fmt"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	c, err := cid.Decode("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Got CID: ", c)

	// Create a cid manually by specifying the 'prefix' parameters
	pref := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   mh.SHA3_256,
		MhLength: -1, // default length
	}

	// And then feed it some data
	c, err = pref.Sum([]byte("Hello World!"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Created CID: ", c)
}
