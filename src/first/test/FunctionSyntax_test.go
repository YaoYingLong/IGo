package test

import (
	"first/function"
	"testing"
)

func TestBaseFunc(t *testing.T) {
	function.BaseFunc()
}

func BenchmarkBaseFunc(t *testing.B) {
	for i := 0; i < t.N; i++ {
		function.BaseFunc()
	}
}
