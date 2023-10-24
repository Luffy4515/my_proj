package sorting

type SortRequest struct {
	Arr []int `json:"arr"`
}

type SortedResponse struct {
	SortedArr []int `json:"sortedarr"`
}

type StrRequest struct {
	Str string `json:"str"`
}

type RepResponse struct {
	RepMap         map[string]int `json:"repetitions map"`
	HighestElement string         `json:"highest element"`
	HiCount        int            `json:"count"`
}
