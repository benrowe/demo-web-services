package entities

type Department struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Manager *Employee `json:"manager"`
}
