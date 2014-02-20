package vindinium

const (
	PLAYER1 = iota
	PLAYER2
	PLAYER3
	PLAYER4
)

type Game struct {
	State       *State
	Board       *Board  `json:"board"`
	Heroes      []*Hero `json:"heroes"`
	Id          string  `json:"id"`
	Finished    bool    `json:"finished"`
	Turn        int     `json:"turn"`
	MaxTurns    int     `json:"maxTurns"`
	Hero        *Hero   `json:"hero"`
	Token       string  `json:"token"`
	Crashed     bool    `json:"crashed"`
	ViewUrl     string  `json:"viewUrl"`
	PlayUrl     string  `json:"PlayUrl"`
	MinesLocs   map[Position]int
	HeroesLocs  map[Position]int
	TavernsLocs map[Position]struct{}
}

func NewGame(state *State) (game *Game) {
	game = state.Game
	game.State = state
	game.Board.parseTiles()

	return
}
