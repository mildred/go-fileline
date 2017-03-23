package fileline

import (
	"regexp"
	"testing"
)

func TestModificationInsert1(t *testing.T) {
	m := Modification{
		SelectRegion: &SelectRegion{
			Start:       regexp.MustCompile("begin"),
			End:         regexp.MustCompile("end"),
			EndIncluded: true,
		},
		InsertLines: []InsertLine{
			InsertLine{
				Lines:   []string{"foo", "bar\n"},
				Pattern: regexp.MustCompile("nd"),
				Last:    true,
			},
		},
	}

	before := `
a
nd
b
-- begin --
a
b
nd
c
-- end --
y
end
z
`
	expected := `
a
nd
b
-- begin --
a
b
nd
c
-- end --
foo
bar
y
end
z
`
	f := ReadString(before)
	f.ModifySingle(m)
	after := f.String()
	if after != expected {
		t.Fatalf("Not expected: %#s", after)
	}
}

func TestModificationInsert2(t *testing.T) {
	m := Modification{
		SelectRegion: &SelectRegion{
			Start: regexp.MustCompile("begin"),
			End:   regexp.MustCompile("end"),
		},
		InsertLines: []InsertLine{
			InsertLine{
				Lines:   []string{"foo\r\n", "bar\n"},
				Pattern: regexp.MustCompile("nd"),
				Before:  true,
			},
		},
	}

	before := `
a
nd
b
-- begin --
a` + "\r" + `
b
nd
nd
c
-- end --
y
end
z
`
	expected := `
a
nd
b
-- begin --
a` + "\r" + `
b
foo` + "\r" + `
bar
nd
nd
c
-- end --
y
end
z
`
	f := ReadString(before)
	f.ModifySingle(m)
	after := f.String()
	if after != expected {
		t.Fatalf("Not expected: %#s", after)
	}
}

func TestModificationDelete(t *testing.T) {
	m := Modification{
		DeleteLines: []DeleteLine{
			DeleteLine{
				Pattern: regexp.MustCompile("toto"),
			},
		},
	}

	before := `
totoa
b
-- begin --
totoc
d
totoe
-- end --
y
end
totoz
`
	expected := `
b
-- begin --
d
-- end --
y
end
`
	f := ReadString(before)
	f.ModifySingle(m)
	after := f.String()
	if after != expected {
		t.Fatalf("Not expected: %#s", after)
	}
}

func TestModificationReplace(t *testing.T) {
	m := Modification{
		ReplaceLines: []ReplaceLine{
			ReplaceLine{
				Pattern:     regexp.MustCompile("<(.*)>"),
				Replacement: ">$1<",
				All:         true,
			},
		},
	}

	before := `
toto<a>
b
-- begin --
toto<c>
d
toto<e>
-- end --
y
end
toto<z>
`
	expected := `
toto>a<
b
-- begin --
toto>c<
d
toto>e<
-- end --
y
end
toto>z<
`
	f := ReadString(before)
	f.ModifySingle(m)
	after := f.String()
	if after != expected {
		t.Fatalf("Not expected: %#s", after)
	}
}
