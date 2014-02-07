package slug_test

import (
	. "slug"
	"testing"
)

var (
	table = []struct{ input, expected string }{
		{"Hello World!-", "hello-world"},
		{"Ä ä Ö ö Ü ü ß", "a-a-o-o-u-u-ss"},
		{"I ♥ the Gopher", "i-love-the-gopher"},
		{"I ♥ the Gopher ♥", "i-love-the-gopher-love"},
		{"ABC世界def-", "abc-def"},
		{"---", ""},
	}
)

func TestSlug(t *testing.T) {
	for _, s := range table {
		r := Slug(s.input)
		t.Logf("%q -> %q", s.input, r)

		if r != s.expected {
			t.Errorf("Expecting: %q but got %q for %q", s.expected, r, s.input)
		}
	}
}

func BenchmarkSlug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slug("I ♥ the Gopher")
	}
}
