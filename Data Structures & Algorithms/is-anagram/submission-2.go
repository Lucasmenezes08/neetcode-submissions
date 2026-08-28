func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    words := make(map[rune]int)

    for _, str := range(s){
        words[str] += 1
    } 
    
    for _, str := range(t){
        words[str] -= 1
        if words[str] < 0{
            return false
        } 
    }
    return true
}
