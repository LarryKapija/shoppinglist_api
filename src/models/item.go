package models

// State = Pendiente compra | Comprado | Descartado
type State int

const (
	_ State = iota
	Pending
	Bought
	Discarded
)

var toID = map[string]State{
	"pending":   Pending,
	"bought":    Bought,
	"discarded": Discarded,
}

func ToId(state string) State {
	return toID[state]
}

type Item struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	State    State   `json:"state"`
}
