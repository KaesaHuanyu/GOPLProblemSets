package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}

func Popcount(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}

	return int(count)
}
