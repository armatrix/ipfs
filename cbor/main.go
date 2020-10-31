package main

import (
	"fmt"

	cbor "github.com/ipfs/go-ipld-cbor"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	b := []byte("jtgqWCcAAXGg5AIg8DSNUI/zS9dTQE8jfH7VfePoA/QqTT7HeYzMQJJq6GlAQNgqWCcAAXGg5AIgFhh6X3zHQfjVeLWK9459nTaJK4JuzIdQ1Fqb1/3cbDFASgAFNpHNwBlCaYHYKlgnAAFxoOQCIBj+asxho6NrDDc8SjqOpkuBK/LKm1KAUJCceNQIVYoM2CpYJwABcaDkAiCGJ8nlCylT0ueqre6iSE1jtYcHekpX2P7+0lWB5z1udtgqWCcAAXGg5AIgJhBXOcMlxjb8Q176oE2Gt79R2BFCZXzr9vg6a8IWex/YKlgnAAFxoOQCIIxkEi8WgEK+1gGBj6i1Z73nEfN8qHjEZWGOCfUFvckQGgACkn0A2CpYJwABcaDkAiC7zbfzoQqKueJfHqfUj5k3X8Wgg+rrS4koar3nZjhsZUA=")
	// b = []byte("bafyreihtkjnakxpxl6lvx6shs2ee65pq3fc3a3ea5gyw4xjgsfky6znwzi")
	res, err := cbor.Decode(b, mh.SHA2_256, -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	nd, _ := cbor.WrapObject(map[string]string{
		"name": "foskdjfklsdjflksjdflsjdlfjdsklfjslkdjflsdkjflksjdflksjdflksjdflksdjflksjdflksjdflkjsdlfkjslkdjo",
	}, mh.SHA2_256, -1)
	fmt.Println(nd)

	var m map[string]string
	cbor.DecodeInto(nd.RawData(), &m)
	fmt.Println(m)
}
