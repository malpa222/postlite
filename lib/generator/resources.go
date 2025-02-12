package generator

type Operation int

const (
	Copy Operation = iota
	Parse
)

var resourcePaths = map[string]Operation{
	"assets":     Copy,
	"index.html": Copy,
	"pages":      Parse,
	"posts":      Parse,
}
