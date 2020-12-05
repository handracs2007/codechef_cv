// https://www.codechef.com/problems/CV
package main

import "fmt"

func isVowel(chr rune) bool {
    return chr == 'a' || chr == 'e' || chr == 'i' || chr == 'o' || chr == 'u'
}

func isNotVowel(chr rune) bool {
    return !isVowel(chr)
}

func main() {
    var numOfCases int
    var length int
    var count int
    var text string

    _, _ = fmt.Scan(&numOfCases)

    for i := 0; i < numOfCases; i++ {
        count = 0

        _, _ = fmt.Scan(&length)
        _, _ = fmt.Scanln(&text)

        for idx, chr := range text {
            if isVowel(chr) {
                if idx-1 >= 0 {
                    if isNotVowel(rune(text[idx-1])) {
                        count++
                    }
                }
            }
        }

        fmt.Println(count)
    }
}
