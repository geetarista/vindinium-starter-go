package vindinium

type State struct {
	Game    *Game  `json:"game"`
	Hero    *Hero  `json:"hero"`
	Token   string `json:"token"`
	ViewUrl string `json:"viewUrl"`
	PlayUrl string `json:"PlayUrl"`
}
