package calculator

import (
	"math"
)

// write code here

//const PI = 3.14

type Cube struct {
	Length float64
}

type Pyramid struct {
	Length float64
	Weight float64
	Height float64
}

type Sphere struct {
	Radius float64
}

type VolumeCalculator interface {
	CalculateVolumeShape() float64
}

func CalculateVolume(v VolumeCalculator) float64 {
	return v.CalculateVolumeShape()
}

/// CUBE

func (c *Cube) CalculateVolumeShape() float64 {
	return math.Pow(c.Length, 3)
}

func NewCube (length float64) *Cube {
	return &Cube {
		Length: length,
	}
}

/// PYRAMID

func (p *Pyramid) CalculateVolumeShape() float64 {
	return 1 / 3.0 * p.Length * p.Weight * p.Height
}

func NewPyramid (l , w, h float64) *Pyramid {
	return &Pyramid{
		Length: l,
		Weight: w,
		Height: h,
	}
}

/// SPHERE

func (s *Sphere) CalculateVolumeShape() float64 {
	return Round(4 / 3.0 * math.Pi * math.Pow(s.Radius,3) , 2)
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{
		Radius: radius,
	}
}



func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}