package es

import "golang.org/x/exp/slices"

func MinStreamID(ids ...string) (string, bool) {
	if len(ids) == 0 {
		return "", false
	}
	return slices.Min(ids), true
}

func MaxStreamID(ids ...string) (string, bool) {
	if len(ids) == 0 {
		return "", false
	}
	return slices.Max(ids), true
}
