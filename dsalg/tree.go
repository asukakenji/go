package dsalg

type direction *bool

var (
	left  direction = func() direction { b := false; return &b }()
	right direction = func() direction { b := true; return &b }()
)
