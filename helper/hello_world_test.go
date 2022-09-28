package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// table benchmark
func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Budi",
			request: "Budi",
		},
		{
			name:    "Andra",
			request: "Andra",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("tkba", func(b *testing.B) {
		HelloWorld("tkba")
	})
}

// benchmartk
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("a")
	}
}

// test table
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Taufiq",
			request:  "Taufiq",
			expected: "Hello Taufiq",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Bayu",
			request:  "Bayu",
			expected: "Hello Bayu",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Before-after unit test
func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	m.Run() // eksekusi semua unit test
	fmt.Println("After unit test")
}

// Test sub test
func TestSubTest(t *testing.T) {
	t.Run("SubTest", func(t *testing.T) {
		result := HelloWorld("Sub Test")
		assert.Equal(t, "Hello Sub Test", result, "Result must be 'Hello Sub Test'")
	})
	t.Run("SubTest2", func(t *testing.T) {
		result := HelloWorld("Sub Test 2")
		assert.Equal(t, "Hello Sub Test 2", result, "Result must be 'Hello Sub Test'")
	})
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Taufiq Kurniawan")
	if result != "Hello Taufiq Kurniawan" {
		//	unit test failed
		t.Fail()
	}
	fmt.Println("TestHelloWorld Done!")
}

func TestHelloWorldTaufiq(t *testing.T) {
	result := HelloWorld("taufiq")
	if result != "Hello taufiq" {
		//	unit test failed
		t.Error("Result must be 'Hello taufiq'")
	}
	fmt.Println("TestHelloWorldTaufiq Done!")
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("kurniawan")
	if result != "Hello kurniawan" {
		t.Fatal("Result must be 'Hello kurniawan'")
	}
	fmt.Println("TestHelloWorldKurniawan Done!")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Taufiq")
	assert.Equal(t, "Hello Taufiq", result, "Result must be 'Hello Taufiq'")
	fmt.Println("TestHelloWorldAssert Done!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Kurniawan")
	require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	fmt.Println("TestHelloWorld with Require Done!")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipped not running at windows")
	}
	fmt.Println("After skipping")
}
