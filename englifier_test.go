package englifier

import (
	"reflect"
	"testing"
)

func Test_CollectWords(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty html", args{"<html></html>"}, []string{}},
		{"empty body", args{"<html><body></body></html>"}, []string{}},
		{"single word not in body", args{"<html>the</html>"}, []string{}},
		{"single word in head", args{"<html><head>the</head></html>"}, []string{}},
		{"single word in head and body", args{"<html><head>whatever</head><body>the</body></html>"}, []string{"the"}},
		{"two words not in body", args{"<html>the quick</html>"}, []string{}},
		{"two words in body", args{"<html><body>the quick</body></html>"}, []string{"the", "quick"}},
		{"two words with punctuation", args{"<html><body>the, quick.</body></html>"}, []string{"the", "quick"}},
		{"with more tags", args{"<html><body>the <b>quick</b></body></html>"}, []string{"the", "quick"}},
		{"with attributes", args{`<html class="one"><body>the <b id="two">quick</b></body></html>`}, []string{"the", "quick"}},
		{"with escape", args{"<html><body>the <b>quick&amp;brown</b></body></html>"}, []string{"the", "quickbrown"}},
		{"single word in body with attribute", args{`<html><body class="apple">the</body></html>`}, []string{"the"}},
		// TODO collect words with number of occurences
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CollectWords(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectWords() = %#v, want %#v", got, tt.want)
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
		b  bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"add empty to empty", args{[]string{}, "", true}, []string{}},
		{"add empty to non-empty", args{[]string{"the"}, "", true}, []string{"the"}},
		{"add non-empty to empty", args{[]string{}, "the", true}, []string{"the"}},
		{"add non-empty to non-empty", args{[]string{"the"}, "quick", true}, []string{"the", "quick"}},
		{"add to longer", args{[]string{"the", "quick"}, "brown", true}, []string{"the", "quick", "brown"}},
		//{"add repeated", args{[]string{"the", "quick"}, "the"}, []string{"the", "quick"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addIfNonEmpty(tt.args.ws, tt.args.w, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addIfNonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
