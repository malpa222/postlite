package blogfsys

type FileKind int

const (
	MD FileKind = 1 << iota
	HTML
	CSS
	YAML
	Media
	Dir
	All = MD | HTML | CSS | YAML | Media | Dir
)
