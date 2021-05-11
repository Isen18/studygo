package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	in := "abc"
	sep := "b"
	got := Split(in, sep)
	want := []string{"a", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %#v, but got %#v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	type testCase struct {
		in   string
		sep  string
		want []string
	}

	testGroup := []testCase{
		{"abc", "b", []string{"a", "c"}},
		{"abc", "c", []string{"ab", ""}},
	}

	for _, tc := range testGroup {
		got := Split(tc.in, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("want %#v, but got %#v", tc.want, got)
		}
	}
}

func TestSplit3(t *testing.T) {
	type testCase struct {
		in   string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": {"abc", "b", []string{"a", "c"}},
		// "case_2": {"abc", "bb", []string{"", "a", "b2", "c", ""}},
		"case_3": {"abc", "c", []string{"ab", ""}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.in, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %#v, but got %#v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
