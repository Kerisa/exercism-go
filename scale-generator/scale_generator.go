package scale

import "strings"
// var majorType = []string { "C", "G", "D", "A", "E", "B", "F#", "F", "Bb", "Eb", "Ab", "Db", "Gb" }
// var minorType = []string { "a", "e", "b", "f#", "c#", "g#", "d#", "d", "g", "c", "f", "bb", "eb" }

var chromaticScale = []string { "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#" }
var flatChromaticScale = []string { "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab" }
var step = map[rune]int { 'm':1, 'M':2, 'A':3 }

// Scale doc here
func Scale(tonic, interval string) []string {
	var useType []string
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		useType = flatChromaticScale
	default:
		useType = chromaticScale
	}

	var startPos int
	for i,str := range useType {
		if strings.ToLower(tonic) == strings.ToLower(str) {
			startPos = i
			break
		}
	}

	if len(interval) == 0 {
		interval = strings.Repeat("m", len(useType))
	}

	result := []string{}
	for _,inv := range interval {
		st := step[inv]
		result = append(result, useType[startPos])
		startPos = (startPos + st) % len(useType)
	}
	return result
}