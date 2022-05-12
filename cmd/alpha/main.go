package main

import (
	"log"

	"gopkg.in/yaml.v2"

	"github.com/xiejw/alpha/nn"
)

var data = `
kind: network
version: single_tower.v1
batch_size: &bs 32
placeholders:
  - kind: input
    name: x
    shape: [*bs, 16]
  - kind: output
    name: o
    shape: [*bs, 16]
  - kind: label
    name: y
    shape: [*bs, 1]
layers:
- kind: layer
  version: dense.v1
  inputs: [x]
  properties:
    hidden_dim: 64
- kind: layer
  version: dense.v1
  properties:
    hidden_dim: 128
  outputs: [y]
`

func main() {
	// ---------------------------------------------------------------------
	// Print the original data
	log.Printf("--- data original:\n%v\n\n", data)

	// ---------------------------------------------------------------------
	// Print the unmarshal result
	n := nn.Network{}
	err := yaml.Unmarshal([]byte(data), &n)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- n:\n%v\n\n", n)

	// ---------------------------------------------------------------------
	// Print the marshal result
	d, err := yaml.Marshal(&n)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- t dump:\n%s\n\n", string(d))
}
