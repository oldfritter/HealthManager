package models

var AllComponents = make(map[string]Component)

type Component struct {
	Name       string      `json:"name"`
	Kind       string      `json:"kind"`
	Subscribed string      `json:"subscribed"`
	Details    interface{} `json:"details"`
	Timestamp  int64       `json:"timestamp"`
}
