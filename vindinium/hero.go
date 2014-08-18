package vindinium

type Hero struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    string    `json:"userId"`
	Elo       int       `json:"elo"`
	Pos       *Position `json:"pos"`
	Life      int       `json:"life"`
	Gold      int       `json:"gold"`
	MineCount int       `json:"mineCount"`
	SpawnPos  *Position `json:"spawnPos"`
	Crashed   bool      `json:"crashed"`
}
