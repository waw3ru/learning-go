package iteration

import "strings"

/**
 * Using strings.Builder
 */
func Repeat(char string, n int) string {
	var r strings.Builder

	for range n {
		r.WriteString(char)
	}

	return r.String()
}

/**
 * Without using strings.Builder
 */
func RepeatNoStrBuilder(char string, n int) string {
	var r string

	for range n {
		r = r + char
	}

	return r
}

/*
Benchmarks Report
------------------------------------------------

=== RUN   BenchmarkRepeat
BenchmarkRepeat
BenchmarkRepeat-8       12945216                82.59 ns/op            8 B/op          1 allocs/op
PASS
ok      github.com/waw3ru/learning-go/iteration 1.430s

=== RUN   BenchmarkRepeatNoStrBuilder
BenchmarkRepeatNoStrBuilder
BenchmarkRepeatNoStrBuilder-8            3725596               316.8 ns/op            16 B/op          4 allocs/op
PASS
ok      github.com/waw3ru/learning-go/iteration 1.518s


-- StringBuilder is faster and uses less memory
*/
