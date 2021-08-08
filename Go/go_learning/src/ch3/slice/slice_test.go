package slice

import "testing"

func TestSliceInit(t *testing.T) {
	// var s0 []int
	// t.Log(len(s0), cap(s0))
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Noc", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]

	// 很有意思！
	t.Log(summer, len(summer), cap(summer))
	summer = append(summer, "test")
	summer[0] = "Unknown"

	t.Log(summer)
	t.Log(Q2)
	t.Log(year)
}
