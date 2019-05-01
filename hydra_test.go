package hydra_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Warashi/hydra"
)

type Struct struct {
	String string
	Int    int
	Nested Nested
}

type Nested struct {
	String string
	Int    int
}

func TestJSONLoader(t *testing.T) {
	wanted := Struct{
		String: "string",
		Int:    1,
		Nested: Nested{
			String: "nested_string",
			Int:    2,
		},
	}

	var s Struct
	if err := hydra.JSONLoader("testdata/data.json").Load(&s); err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(s, wanted) {
		t.Errorf(
			"JSONLoader() = %+v, wanted = %+v, diff = %+v",
			s,
			wanted,
			cmp.Diff(s, wanted),
		)
	}
}
