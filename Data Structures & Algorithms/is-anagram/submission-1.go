import "maps"

func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    words := make(map[rune]int)
    words2 := make(map[rune]int) 

    for _, str := range(s){
        words[str] += 1
    } 
    for _, str2 := range(t){
        words2[str2] += 1
    }

    if maps.Equal(words, words2){
        return true
    }
    return false
}
