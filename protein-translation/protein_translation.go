
package protein

import "errors"

var STOP = errors.New("stop")
var ErrInvalidBase = errors.New("errInvalidBase")

var codeonMap = map[string]string{
	"AUG" : "Methionine",
	"UUU" : "Phenylalanine",
	"UUC" : "Phenylalanine",
	"UUA" : "Leucine",
	"UUG" : "Leucine",
	"UCU" : "Serine",
	"UCC" : "Serine",
	"UCA" : "Serine",
	"UCG" : "Serine",
	"UAU" : "Tyrosine",
	"UAC" : "Tyrosine",
	"UGU" : "Cysteine",
	"UGC" : "Cysteine",
	"UGG" : "Tryptophan",
	"UAA" : "STOP",
	"UAG" : "STOP",
	"UGA" : "STOP",
}

// FromCodon doc here
func FromCodon(in string) (string, error) {
	if p,ok := codeonMap[in]; ok {
		if p == "STOP" {
			return "", STOP
		}
		return p, nil
	}
	return "", ErrInvalidBase
}

// FromRNA doc here
func FromRNA(in string) ([]string, error) {
	const CODONLENGTH int = 3
	var result = []string{}
	for i := 0; i < len(in) - CODONLENGTH + 1; i += CODONLENGTH {
		switch p,ok := FromCodon(in[i : i+CODONLENGTH]); ok {
		case nil:
			result = append(result, p)

		case ErrInvalidBase:
			return result, ok

		case STOP:
			return result, nil
		}
	}
	return result, nil
}
