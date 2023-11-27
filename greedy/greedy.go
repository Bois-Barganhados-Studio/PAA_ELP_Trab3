package greedy

func GetJobs(deadlines []int, penalties []int) int {

	penalty := 0
	hash := make(map[int]int)

	for i := 0; i < len(penalties); i++ {
		if hash[deadlines[i]] == 0 {
			hash[deadlines[i]] = penalties[i]
			penalty += penalties[i]
		} else if hash[deadlines[i]] < penalties[i] {
			penalty -= hash[deadlines[i]]
			penalty += penalties[i]
			hash[deadlines[i]] = penalties[i]
		}

	}

	return penalty
}
