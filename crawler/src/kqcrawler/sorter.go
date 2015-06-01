package kqcrawler

//****** map sorter
type PlayerScoreSorter []PlayerScore

func NewPlayerScoreSorter(m map[int]PlayerScore) PlayerScoreSorter {
	ms := make(PlayerScoreSorter, 0, len(m))
	for _, v := range m {
		ms = append(ms, v)
	}

	return ms
}

func (ms PlayerScoreSorter) Len() int {
	return len(ms)
}

func (ms PlayerScoreSorter) Less(i, j int) bool {
	if ms[i].Matched {
		return true
	}

	return ms[i].Score > ms[j].Score
}

func (ms PlayerScoreSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
