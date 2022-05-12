package nn

type Network struct {
	// Required by Component
	Kind    string
	Version string

	// Optional. Wil be overwritten by Component.Sign()
	Checksum string

	// Required. The batch size of this network.
	BatchSize int `yaml:"batch_size"`

	// Required. The placeholders for inputs, outputs, labels, etc.
	PlaceHolders []PlaceHolder

	// Required.
	//
	// The details requirement is defined by the underlying implementaiton.
	//
	// For example, for single tower network:
	//   - The first layer must have inputs. Othe layers are allowed to take
	//     inputs as well, but optional.
	//   - The last layer must have outputs. No other layers are allowed to
	//     have outputs.
	Layers []Layer // Required
}

type PlaceHolder struct {
	Kind  string // Required: input, output, label
	Name  string // Required. Unique in the scope (e.g., Network).
	Shape []int  `yaml:",flow"` // Required
}

type Layer struct {
	// Required by Component
	Kind    string
	Version string

	// Optional. Wil be overwritten Component.Sign()
	Checksum string

	Input      []string               // Optional
	Outputs    []string               // Optional
	Properties map[string]interface{} // Optional
}
