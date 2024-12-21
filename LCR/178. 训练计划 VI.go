func trainingPlan(actions []int) int {
    m := map[int]int{}
    for _, value := range actions {
        if _, ok := m[value]; ok {
            m[value] += 1
        } else {
            m[value] = 1
        }
    }
    for key, value := range m {
        if value == 1 {
            return key
        }
    }
    return -1
}
