package protein

import (
	"errors"
)

// ErrStop error for stop
var ErrStop = errors.New("stop")

// ErrInvalidBase if the base is invalid
var ErrInvalidBase = errors.New("invalid base")

var proteines = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon translates codons to Proteins
func FromCodon(input string) (string, error) {
	protein, existing := proteines[input]

	if !existing {
		return "", ErrInvalidBase
	}

	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}

// FromRNA converts RNA to proteins
func FromRNA(input string) ([]string, error) {
	set := make(map[string]struct{}) // New empty set
	proteinList := []string{}

	for i := 0; i <= len(input)-3; i += 3 {
		protein, err := FromCodon(input[i : i+3])

		if err == ErrStop {
			break
		}

		if err == ErrInvalidBase {
			return proteinList, ErrInvalidBase
		}

		_, ok := set[protein]
		if ok {
			continue
		}

		proteinList = append(proteinList, protein)
		set[protein] = struct{}{}
	}

	return proteinList, nil
}
