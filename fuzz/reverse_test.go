package main

import (
	"testing"
	"unicode/utf8"
)

// go test -run=FuzzReverse
// go test -fuzz=FuzzReverse -fuzztime=10s
// go test -fuzz=FuzzReverse -fuzztime=10s -fuzzseed=12345
// go test -fuzz=FuzzReverse -fuzztime=10s -fuzzseed=12345 -fuzzcount=1000
// go test -fuzz=FuzzReverse -fuzztime=10s -fuzzseed=12345 -fuzzcount=1000 -fuzzworker=4
// go test -fuzz=FuzzReverse -fuzztime=10s -fuzzseed=12345 -fuzzcount=1000 -fuzzworker=4 -fuzzfunc=Reverse

// func FuzzReverse(f *testing.F) {
//     testcases := []string{"Hello, world", " ", "!12345"}
//     for _, tc := range testcases {
//         f.Add(tc)  // Use f.Add to provide a seed corpus
//     }

//     f.Fuzz(func(t *testing.T, orig string) {
//         rev := Reverse(orig)
//         doubleRev := Reverse(rev)
//         t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
//         if orig != doubleRev {
//             t.Errorf("Before: %q, after: %q", orig, doubleRev)
//         }
//         if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//             t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//         }
//     })

//     // f.Fuzz(func(t *testing.T, orig string) {
//     //     rev := Reverse(orig)
//     //     doubleRev := Reverse(rev)
//     //     if orig != doubleRev {
//     //         t.Errorf("Before: %q, after: %q", orig, doubleRev)
//     //     }
//     //     if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//     //         t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//     //     }
//     // })
// }

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
//                 t.Errorf("Reverse: %q, want %q", rev, tc.want)
//         }
//     }
// }

func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}