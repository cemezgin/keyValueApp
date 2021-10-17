package key_value

type Items struct {
	Items []Item
}

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}