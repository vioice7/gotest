// this package should be in the GOPATH env variable
package t10

// the sum of all numbers passed as variatic parameter
// also it will return an int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
