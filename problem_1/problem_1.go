package problem_1

//
// https://leetcode.com/problems/two-sum/description/
//

func Main(input []int, target int) []int {
	output := []int{}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == target {
				output = append(output, i)
				output = append(output, j)
			}
		}
	}
	m := map[int]int{}
	for _, out := range output {
		m[out] = out
	}
	output = []int{}
	for k := range m {
		output = append(output, k)
	}
	return output
}
