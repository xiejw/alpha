package nn

type Network struct {
	Kind         string        // required
	Version      string        // required
	BatchSize    int           `yaml:"batch_size"` // required
	PlaceHolders []PlaceHolder // required
	Layers       []Layer       // required
	Checksum     string        // optional. will overwrite anyway
}

type PlaceHolder struct {
	Name  string // required
	Shape []int  `yaml:",flow"` // required
	Kind  string // required: input, output, label
}

type Layer struct {
	Kind       string                 // required
	Version    string                 // required
	Input      []string               // optional
	Outputs    []string               // optional
	Properties map[string]interface{} // optional
	Checksum   string                 // optional. will overwrite anyway
}
