package officialdoc

/** 闭包 官方示例 **/
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/** 斐波那契数列 **/
func Fibonacci() func(int) int {
	return func(x int) int {
		sum := 0
		for i := 0; i <= x; i++ {
			sum += i
		}
		return sum
	}
}
