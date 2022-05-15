package nn

type Component interface {
	Kind() string    // The type, e.g., network, layer, etc.
	Version() string // The name with version, dense.v1. Unique under type.

	// Fill default values for all required fields. May mutabe some fields
	// if needed.
	AdmissionReivew() error

	// Will also infer shape and input names.
	InferOutputs() error

	// Validate all required fields and optional fields if present. Will
	// assume AdmissionReivew is called.
	Validate() error

	// Sign the all required fields (input shapes, output shapes, etc) and
	// generate checksum. Will assume Validate is called.
	Sign() error
}
