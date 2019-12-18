package _struct

type Data []struct {
	One []struct {
		Ts    int `json:"ts"`
		Value int `json:"value"`
	} `json:"one"`
}

func NewData() Data {
	res := make([]struct {
		One []struct {
			Ts    int `json:"ts"`
			Value int `json:"value"`
		}
	}, 0)
	return Data(res)
}
