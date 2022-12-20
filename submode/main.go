package main


type analyze struct {

}

func (a *analyze) GetAnalyze() []int {
	return []int{1, 2, 3, 4}
}

var manager analyze
