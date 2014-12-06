// Copyright 2014 Rob Pike. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Kata accepts as argument text romaji
text. It copies the text to standard output, converting
the text that matches the romaji format to katakana.
Note that the output may not be accurate because of
false matches and the inability to generate things like
the tsu consonant-extending symbol.
*/
package main // import "robpike.io/cmd/kata"

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	do(strings.Join(os.Args[1:], " "))
}

func do(s string) {
	wid := 0
	for i := 0; i < len(s); i += wid {
		if i+3 <= len(s) {
			cha, ok := three[s[i:i+3]]
			if ok {
				wid = 3
				fmt.Printf("%s", cha)
				continue
			}
		}
		if i+2 <= len(s) {
			ka, ok := two[s[i:i+2]]
			if ok {
				wid = 2
				fmt.Printf("%s", ka)
				continue
			}
		}
		wid = 1
		a, ok := one[s[i:i+1]]
		if ok {
			fmt.Printf("%s", a)
			continue
		}
		fmt.Printf("%s", s[i:i+1])
	}
	fmt.Print("\n")
}

var one = map[string]string{
	"a": "ア",
	"i": "イ",
	"u": "ウ",
	"e": "エ",
	"o": "オ",
	"n": "ン",
}

var two = map[string]string{
	"ka": "カ",
	"ga": "ガ",
	"ki": "キ",
	"gi": "ギ",
	"ku": "ク",
	"gu": "グ",
	"ke": "ケ",
	"ge": "ゲ",
	"ko": "コ",
	"go": "ゴ",
	"sa": "サ",
	"za": "ザ",
	"si": "シ",
	"zi": "ジ",
	"su": "ス",
	"zu": "ズ",
	"se": "セ",
	"ze": "ゼ",
	"so": "ソ",
	"zo": "ゾ",
	"ta": "タ",
	"da": "ダ",
	"ti": "チ",
	"di": "ヂ",
	"tu": "ツ",
	"du": "ヅ",
	"te": "テ",
	"de": "デ",
	"to": "ト",
	"do": "ド",
	"na": "ナ",
	"ni": "ニ",
	"nu": "ヌ",
	"ne": "ネ",
	"no": "ノ",
	"ha": "ハ",
	"va": "バ",
	"pa": "パ",
	"hi": "ヒ",
	"vi": "ビ",
	"pi": "ピ",
	"fu": "フ",
	"bu": "ブ",
	"pu": "プ",
	"he": "ヘ",
	"ve": "ベ",
	"pe": "ペ",
	"ho": "ホ",
	"vo": "ボ",
	"po": "ポ",
	"ma": "マ",
	"mi": "ミ",
	"mu": "ム",
	"me": "メ",
	"mo": "モ",
	"ya": "ヤ",
	"yu": "ユ",
	"yo": "ヨ",
	"ra": "ラ",
	"ri": "リ",
	"ru": "ル",
	"re": "レ",
	"ro": "ロ",
	"wa": "ワ",
	"wi": "ヰ",
	"we": "ヱ",
	"wo": "ヲ",
	"vu": "ヴ",
}

var three = map[string]string{
	"shi": "シ",
	"chi": "チ",
	"tsu": "ツ",
	"cha": "チャ", // TODO IS THIS RIGHT
	"chu": "チュ", // TODO IS THIS RIGHT
	"cho": "チョ", // TODO IS THIS RIGHT
	// NEED MORE
}
