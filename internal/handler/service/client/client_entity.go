package client

const (
	CoorX       string = "coorX"
	CoorY       string = "coorY"
	Orientation string = "orientation"
	Instruction string = "instruction"
	Turn        string = "turn"
	Result      string = "result"
)

const (
	N string = "N"
	S string = "S"
	E string = "E"
	W string = "W"
)

type Request struct {
	CoorX       int
	CoorY       int
	Orientation string
	Instruction string
}

type Response struct {
	CoorX       int
	CoorY       int
	Orientation string
}
