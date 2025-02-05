package generator

import "fmt"

type ReadError struct {
	context string
	err     error
}

func (r *ReadError) Error() string {
	return fmt.Sprintf("Error reading %s: %s", r.context, r.err)
}

type WriteError struct {
	context string
	err     error
}

func (w *WriteError) Error() string {
	return fmt.Sprintf("Error writing %s: %s", w.context, w.err)
}

type CopyError struct {
	context string
	err     error
}

func (c *CopyError) Error() string {
	return fmt.Sprintf("Error copying %s: %s", c.context, c.err)
}

type DeleteError struct {
	context string
	err     error
}

func (d *DeleteError) Error() string {
	return fmt.Sprintf("Error removing %s: %s", d.context, d.err)
}
