package main


func TestSum(t *testing.T) {
	total := Sum(10, 20)

	if total != 30 {
		t.Errorf("Sum(10, 20) = %d, 30 was expected", total)
	} 
}