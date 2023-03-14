package govecshttp

import (
	"math/rand"

	"encoding/json"

	"github.com/cantoramann/govecs"
)

type HttpPoint struct {
	Id           int    `json:"id"`
	Name         string `json:"name,omitempty"`
	govecs.Point `json:"point"`
}

// NewHttpPoint creates a new HttpPoint with the given id, name and coordinates

func NewHttpPoint(id int, name string, x float64, y float64, z float64) *HttpPoint {
	point := govecs.NewPoint(x, y, z)
	return &HttpPoint{id, name, *point}
}

// newRandomPointData creates a new random point - not open to the outside world

func newRandomPointData() *govecs.Point {
	return govecs.NewPoint(rand.Float64(), rand.Float64(), rand.Float64())
}

// NewRandomHttpPoint creates a new HttpPoint with random id, name and coordinates

func NewRandomHttpPoint() *HttpPoint {
	random_point := newRandomPointData()
	return NewHttpPoint(rand.Int(), "unnamed-point", random_point.GetX(), random_point.GetY(), random_point.GetZ())
}

// Encode encodes a HttpPoint into a byte array

func Encode(point *HttpPoint) ([]byte, error) {
	return json.Marshal(point)
}

// Decode decodes a byte array into a HttpPoint

func Decode(data []byte) (*HttpPoint, error) {
	var point HttpPoint
	err := json.Unmarshal(data, &point)
	return &point, err
}
