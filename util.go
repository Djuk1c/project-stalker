package main

import "math"

func ToRadian(i float64) float64 {
	return i * (math.Pi / 180)
}

func ToDegree(i float64) float64 {
	return i * (180 / math.Pi)
}

func GetAngleBetween(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	angle := -90 + math.Atan2(y1-y2, x1-x2)*(180/math.Pi)
	if angle >= 0 {
		return angle
	} else {
		return 360 + angle
	}
}
