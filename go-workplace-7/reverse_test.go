package main

import (
    "testing"
    "unicode/utf8"
)

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
        if err1 != nil {
			t.Skip() // return文の代わりに「t.Skip()」も使える
            //return
        }
		doubleRev, err2 := Reverse(rev)
        if err2 != nil {
			t.Skip()
            //return
        }
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// func TestReverse(t *testing.T) {
//     testcases := []struct {
//         in, want string
//     }{
//         {"Hello, world", "dlrow ,olleH"},
//         {" ", " "},
//         {"!12345", "54321!"},
//     }
//     for _, tc := range testcases {
//         rev := Reverse(tc.in)
//         if rev != tc.want {
//             t.Errorf("Reverse: %q, want %q", rev, tc.want)
//         }
//     }
// }

// 失敗しなければ30秒以内終了する
// go test -fuzz=Fuzz -fuzztime 30s