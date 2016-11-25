package andrea_test

import (
   "github.com/CarolinaRomero33/carolina"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

// Ensure the parser can parse strings into Statement ASTs.

func TestParser_ParseStatement(t *testing.T) {

	//Erasmo
	f, _ := ioutil.ReadFile("")
	z := string(f[:])
	//Bonifaz
	//f, _ := ioutil.ReadFile("C:\\porno\\file2.txt"")
	//z := string(f[:])
	//z:=strings.Fields(r)
	//var tests1 = []struct {
	//	j    string
	//	stmt *andrea.zifstatement
	//	err  string
	//}{}
	var tests = []struct {
		s    string
		stmt *andrea.SelectStatement
		err  string
	}{
		{
			s:    z,
			stmt: &andrea.SelectStatement{
			//Fields: []string{""},
			//TableName: "andrea",
			},
		},
	}
	for i, tt := range tests {
		stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}

	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsezif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsezif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsezswitch()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q:\ error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).ParseSum()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).ParseSumNom()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).ParsezFor()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsezelseif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//

	//for i, tt := range tests1 {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.j)).Parsezif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
