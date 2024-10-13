package model

type Operation struct {
	Op   string    "json:'op'"
	Nums []float64 "json:'nums'"
}
