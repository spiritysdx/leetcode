func hasSameString(s1,s2 string) bool {
    for _, s := range s1 {
        if strings.Index(s2, string(s)) != -1 {
            return true
        }
    }
    return false
}

func maxProduct(words []string) int {
    res := 0
    for i, s1 := range words {
        for _, s2 := range words[i+1:] {
            if hasSameString(s1, s2) {
                continue
            } else {
                tp := len(s1)*len(s2)
                if tp > res {
                    res = tp
                }
            }
        } 
    }
    return res
}
