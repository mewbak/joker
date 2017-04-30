package string

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	. "github.com/candid82/joker/core"
)

var newLine *regexp.Regexp

func padRight(s, pad string, n int) string {
	for {
		s += pad
		if len(s) > n {
			return s[0:n]
		}
	}
}

func padLeft(s, pad string, n int) string {
	for {
		s = pad + s
		if len(s) > n {
			return s[len(s)-n:]
		}
	}
}

func split(s string, r *regexp.Regexp) Object {
	indexes := r.FindAllStringIndex(s, -1)
	lastStart := 0
	result := EmptyVector
	for _, el := range indexes {
		result = result.Conjoin(String{S: s[lastStart:el[0]]})
		lastStart = el[1]
	}
	result = result.Conjoin(String{S: s[lastStart:]})
	return result
}

func join(sep string, seqable Seqable) string {
	seq := seqable.Seq()
	var b bytes.Buffer
	for !seq.IsEmpty() {
		b.WriteString(seq.First().ToString(false))
		seq = seq.Rest()
		if !seq.IsEmpty() {
			b.WriteString(sep)
		}
	}
	return b.String()
}

func isBlank(s Object) bool {
	if s.Equals(NIL) {
		return true
	}
	str := AssertString(s, "").S
	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func capitalize(s string) string {
	if len(s) < 2 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

func escape(s string, cmap Callable) string {
	var b bytes.Buffer
	for _, r := range s {
		if obj := cmap.Call([]Object{Char{Ch: r}}); !obj.Equals(NIL) {
			b.WriteString(obj.ToString(false))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func indexOf(s string, value Object, from int) Object {
	var res int
	if from != 0 {
		s = string([]rune(s)[from:])
	}
	switch value := value.(type) {
	case Char:
		res = strings.IndexRune(s, value.Ch)
	case String:
		res = strings.Index(s, value.S)
	default:
		panic(RT.NewArgTypeError(1, value, "String or Char"))
	}
	if res == -1 {
		return NIL
	}
	return MakeInt(utf8.RuneCountInString(s[:res]) + from)
}

func lastIndexOf(s string, value Object, from int) Object {
	var res int
	if from != 0 {
		s = string([]rune(s)[:from])
	}
	switch value := value.(type) {
	case Char:
		res = strings.LastIndex(s, string(value.Ch))
	case String:
		res = strings.LastIndex(s, value.S)
	default:
		panic(RT.NewArgTypeError(1, value, "String or Char"))
	}
	if res == -1 {
		return NIL
	}
	return MakeInt(utf8.RuneCountInString(s[:res]))
}

func init() {
	newLine, _ = regexp.Compile("\r?\n")
}