package calculator

//private
var add = func(x, y int) int {
	return x + y
}

//public
func AddClient(x, y int) int {
	return add(x, y)
}

func GETADDER() func(int, int) int {
	whoAmI()
	return add
}
