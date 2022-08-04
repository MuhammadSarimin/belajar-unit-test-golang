package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		req  string
	}{
		{name: "Muhammad", req: "Muhammad"},
		{name: "Sarimin", req: "Sarimin"},
		{name: "Imin", req: "Imind"},
		{name: "MuhammadSarimin", req: "Muhammad Sarimin"},
		{name: "MuhammadSariminMuhammad", req: "Muhammad sarmin Muhammad"},
		{name: "MuhammadSariminMuhammadSarimin", req: "Muhammad Sarimin Muhammad Sarimin"},
	}

	for _, be := range benchmarks {
		b.Run(be.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(be.req)
			}
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ahmad")
	}
}
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhammad")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Muhammad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Muhammad")
		}
	})
	b.Run("Sarimin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sarimin")
		}
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		req  string
		exp  string
	}{
		{name: "Muhammad Sarimin", req: "Muhammad Sarimin", exp: "Hello Muhammad Sarimin"},
		{name: "Sarimin", req: "Sarimin", exp: "Hello Sarimin"},
		{name: "Imin", req: "Imin", exp: "Hello Imin"},
		{name: "IminFailed", req: "Imin Failed", exp: "Hello Imin"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := HelloWorld(test.req)
			require.Equal(t, test.exp, res)
		})
	}
}

func TestSubTest(t *testing.T) { // bisa panggil per sub dengan sintaks go test -v -run=TestSubTest/Imin
	t.Run("Sarimin", func(t *testing.T) {
		res := HelloWorld("Sarimin")
		require.Equal(t, "Hello Sarimin", res, "Result must be 'Hello Sarimin'")
	})
	t.Run("Imin", func(t *testing.T) {
		res := HelloWorld("Imin")
		require.Equal(t, "Hello Imin", res, "Result must be 'Hello Imin'")
	})
	t.Run("IminFail", func(t *testing.T) {
		res := HelloWorld("IminFail")
		require.Equal(t, "Hello Imin", res, "Result must be 'Hello IminFail'")
	})
}

func TestMain(m *testing.M) { // hanya akan jalan di satu peckage
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	res := HelloWorld("Sarimin")

	if res != "Hello Sarimin" {
		t.Error("Result is not Hello Sarimin")
	}
}

func TestHelloWorldMhd(t *testing.T) {
	res := HelloWorld("Mhd")

	if res != "Hello Mhd" {
		t.Fatal("Result is not Hello Mhd")
	}
}

func TestHelloWorldAssert(t *testing.T) { // manggil t.Fail() -- di katakan fail dan melanjutkan kodingan setelah nya
	res := HelloWorld("Sarimin")
	assert.Equal(t, "Hello Sarimin", res)
}
func TestHelloWorldRequire(t *testing.T) { // manggil t.FailNow -- di katakan fail dan berhenti untuk testing sekarang
	res := HelloWorld("Sarimin")
	require.Equal(t, "Hello Sarimin", res)
}
