package ketchup

type NodeDOM struct {
	Element    string       `json:"element"`
	Content    string       `json:"content"`
	Children   []*NodeDOM   `json:"children"`
	Attributes []*Attribute `json:"attributes"`
	Style      *Stylesheet  `json:"style"`
	Parent     *NodeDOM     `json:"-"`
}

type Attribute struct {
	Name  string
	Value string
}

type Stylesheet struct {
	Color           *ColorRGBA
	BackgroundColor *ColorRGBA

	FontSize float64
	Display  string
	Position string

	Width  float64
	Height float64
	Top    float64
	Left   float64
}

type ColorRGBA struct {
	R float64
	G float64
	B float64
	A float64
}
