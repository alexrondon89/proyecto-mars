package service

type Request struct {
	Robots []Robot
}

type Response struct {
	Robots []Robot
}

type Robot struct {
	Instructions string
	Position     Position
}

type Position struct {
	CoorX       int
	CoorY       int
	Orientation string
}
