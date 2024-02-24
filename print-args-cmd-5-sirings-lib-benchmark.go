// how to run test:
// for all tests in package: go test -bench=.
// for certain test: go test -bench=JoinWay
package main

import "testing"

func BenchMarkJoinWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinWay()
	}
}
