package englify

import (
	"reflect"
	"testing"
)

func Test_collectWords(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty string", args{""}, []string{}},
		{"single word", args{"the"}, []string{"the"}},
		{"two words", args{"the quick"}, []string{"the", "quick"}},
		{"four words", args{"the quick brown fox"}, []string{"the", "quick", "brown", "fox"}},
		{"middle punctuation", args{"the, quick"}, []string{"the", "quick"}},
		{"end punctuation", args{"the quick."}, []string{"the", "quick"}},
		{"double punctuation", args{"the, quick."}, []string{"the", "quick"}},
		{"immediate escaping", args{"&amp;"}, []string{}},
		{"empty html", args{"<html></html>"}, []string{}},
		{"html single word", args{"<html>the</html>"}, []string{"the"}},
		{"html two words", args{"<html>the quick</html>"}, []string{"the", "quick"}},
		{"html two words with punctuation", args{"<html>the, quick.</html>"}, []string{"the", "quick"}},
		{"html with more tags", args{"<html>the <b>quick</b></html>"}, []string{"the", "quick"}},
		{"html with attributes", args{`<html class="one">the <b id="two">quick</b></html>`}, []string{"the", "quick"}},
		{"html with escape", args{"<html>the <b>quick&amp;brown</b></html>"}, []string{"the", "quickbrown"}},
		// TODO collect only in html body
		// TODO collect words with number of occurences
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collectWords(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectWords() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_isPunctuation(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"space", args{' '}, true},
		{"comma", args{','}, true},
		{"semicolon", args{';'}, true},
		{"lf", args{'\n'}, true},
		{"cr", args{'\r'}, true},
		{"word-char", args{'a'}, false},
		{"hyphen", args{'-'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPunctuation(tt.args.char); got != tt.want {
				t.Errorf("isPunctuation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addIfNonEmpty(t *testing.T) {
	type args struct {
		ws []string
		w  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"add empty to empty", args{[]string{}, ""}, []string{}},
		{"add empty to non-empty", args{[]string{"the"}, ""}, []string{"the"}},
		{"add non-empty to empty", args{[]string{}, "the"}, []string{"the"}},
		{"add non-empty to non-empty", args{[]string{"the"}, "quick"}, []string{"the", "quick"}},
		{"add to longer", args{[]string{"the", "quick"}, "brown"}, []string{"the", "quick", "brown"}},
		//{"add repeated", args{[]string{"the", "quick"}, "the"}, []string{"the", "quick"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addIfNonEmpty(tt.args.ws, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addIfNonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
