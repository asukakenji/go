package dsalg

type flags uint8

const (
	DirNone     flags = 0x0
	DirLeft     flags = 0x1
	DirRight    flags = 0x2
	NodeCreated flags = 0x10
)
