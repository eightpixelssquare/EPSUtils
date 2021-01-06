package epsslice

// String is a slice of strings //////////////////////////////////////////////
type String []string

// IndexOf returns the index of the first occurance of s or -1 if not found
func (g String) IndexOf(s string) int {
	for i, v := range g {
		if v == s {
			return i
		}
	}
	return -1
}

// Copy returns a copy of the slice
func (g String) Copy() String {
	s := make(String, len(g))
	copy(s, g)
	return s
}

// RemoveIndex removes an element at index i
func (g *String) RemoveIndex(i int) bool {
	*g = append((*g)[:i], (*g)[i+1:]...)
	return true
}

// Add adds a string to the slice
func (g *String) Add(n string) {
	*g = append(*g, n)
}

// Remove removes the first occurrence of s
func (g *String) Remove(s string) bool {
	i := g.IndexOf(s)
	if i == -1 {
		return false
	}
	g.RemoveIndex(i)
	return true
}

// StringSliceIndexOf does something
func StringSliceIndexOf(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}

// // ToFIDSlice converts a string slice to a FIDSlice
// func (g String) ToFIDSlice() FIDSlice {
// 	out := FIDSlice{}
// 	for _, v := range g {
// 		out = append(out, FID(v))
// 	}
// 	return out
// }
