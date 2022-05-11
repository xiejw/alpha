package nn

type PlaceHolder struct {
	Name  string
	Shape []int  `yaml:",flow"`
	Kind  string // input, output, label
}

type Layer struct {
	Kind       string
	Version    string
	Input      []string
	Outputs    []string
	Properties map[string]interface{}
}

type Network struct {
	Kind         string
	Version      string
	BatchSize    int `yaml:"batch_size"`
	PlaceHolders []PlaceHolder
	Layers       []Layer
}
