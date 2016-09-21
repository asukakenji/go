package dsalg

type direction *bool

var (
	left  = func() direction { b := false; return &b }()
	right = func() direction { b := true; return &b }()
)
