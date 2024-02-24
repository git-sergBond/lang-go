// how to run test:
// add cmd arguments to program/test
// cd internal
// for all tests in package: go test -bench=.
// WINDOWS: for all tests in package: go test --bench=.
// for certain test: go test -bench=JoinWay
// WINDOWS: for all tests in package: go test --bench=JoinWay
// WINDOWS: for all tests in package: go test --bench=StringBuilderWay
// WINDOWS: for all tests in package: go test --bench=StringBuilderFprintfWay
// BenchmarkJoinWay 2.36 sec
// BenchmarkStringBuilderWay 1.84 sec
// BenchmarkStringBuilderFprintfWay 1.92 sec
package print_args_cmd

import "testing"

func BenchmarkJoinWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinWay()
	}
}

func BenchmarkStringBuilderWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilderWay()
	}
}

func BenchmarkStringBuilderFprintfWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilderFprintfWay()
	}
}
