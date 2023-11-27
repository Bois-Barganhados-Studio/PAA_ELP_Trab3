package branchbound

func GetJobs(deadlines []int, penalties []int) int {
	Sort(deadlines, penalties)
	penalty := 0
	jobsUsed := []int{}
	for i := 0; i < len(deadlines); i++ {
		if i == 0 {
			jobsUsed = append(jobsUsed, i)
			penalty += penalties[i]
			continue
		}
		if deadlines[i] == deadlines[i-1] {
			if penalties[i] > penalties[i-1] {
				jobsUsed[len(jobsUsed)-1] = i
				penalty -= penalties[i-1]
				penalty += penalties[i]
			}
		} else {
			jobsUsed = append(jobsUsed, i)
			penalty += penalties[i]
		}
	}
	return penalty
}

func Sort(deadlines []int, penalties []int) {
	quicksort(deadlines, penalties, 0, len(deadlines)-1)
}

func quicksort(deadlines []int, penalties []int, low, high int) {
	if low < high {
		pivotIndex := partition(deadlines, penalties, low, high)
		quicksort(deadlines, penalties, low, pivotIndex-1)
		quicksort(deadlines, penalties, pivotIndex+1, high)
	}
}

func partition(deadlines []int, penalties []int, low, high int) int {
	pivot := deadlines[high]
	i := low - 1

	for j := low; j < high; j++ {
		if deadlines[j] < pivot {
			i++
			deadlines[i], deadlines[j] = deadlines[j], deadlines[i]
			penalties[i], penalties[j] = penalties[j], penalties[i]
		}
	}

	deadlines[i+1], deadlines[high] = deadlines[high], deadlines[i+1]
	penalties[i+1], penalties[high] = penalties[high], penalties[i+1]

	return i + 1
}
