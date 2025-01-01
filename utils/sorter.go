package utils

//* currently unused

type intsx16 struct {
	i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16 int
}
type intsx6 struct {
	i1, i2, i3, i4, i5, i6 int
}

//Bytesum sums individual bytes (cast to ints) in given []byte
func Bytesum(b []byte) int {
	sum := 0
	for _, v := range b {
		sum += int(v)
	}
	return sum
}
