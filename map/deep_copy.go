package _map

import (
	"encoding/json"

	"github.com/mohae/deepcopy"
)

type State struct {
	Name   string
	Status string
	Done   int64
	Total  int64
}

func DeepCopy(m map[string]State) map[string]State {
	return deepcopy.Copy(m).(map[string]State)
}

func ManualCopy(m map[string]State) map[string]State {
	res := make(map[string]State, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

func MarshallCopy(m map[string]State) map[string]State {
	s, _ := json.Marshal(m)
	res := make(map[string]State)
	_ = json.Unmarshal(s, &res)
	return res
}
