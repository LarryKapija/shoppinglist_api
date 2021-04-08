package models

// State = Pendiente compra | Comprado | Descartado
type State int

const (
	Pending State = iota + 1
	Bought
	Discarded
)

type Item struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	State    State   `json:"state"`
}
