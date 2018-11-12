
package triangle

import "sort"
import "math"

type Kind int

const (
    NaT = iota  // not a triangle
    Equ         // equilateral
    Iso         // isosceles
    Sca         // scalene
)

func KindFromSides(a, b, c float64) Kind {
	arr := []float64{a, b, c}
	sort.Float64s(arr)

	// NaN is less than others
	if math.IsNaN(arr[0]) || math.IsInf(arr[0], 0) || math.IsInf(arr[2], 0) || arr[0] <= 0 {
        return NaT
	} else if arr[0] + arr[1] < arr[2] {
        return NaT
	} else if arr[0] == arr[2] {
        return Equ
	} else if arr[0] == arr[1] || arr[1] == arr[2] {
        return Iso
	} else {
        return Sca
    }
}