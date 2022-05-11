package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/xiejw/alpha/nn"
)

var data = `
kind: network
version: v1
batch_size: &bs 32
placeholders:
  - name: x
    shape: [*bs, 16]
    kind: input
  - name: o
    shape: [*bs, 16]
    kind: output
  - name: y
    shape: [*bs, 1]
    kind: label
layers:
- kind: dense
  version: v1
  inputs: [x]
  properties:
    hidden_dim: 64
- kind: dense
  version: v1
  properties:
    hidden_dim: 128
  outputs: [y]
`

func main() {
	t := nn.Network{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
