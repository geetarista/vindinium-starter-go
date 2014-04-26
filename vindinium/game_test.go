package vindinium

import (
	"encoding/json"
	. "gopkg.in/check.v1"
)

var (
	_ = Suite(&GameSuite{})
)

type GameSuite struct {
	game  *Game
	state *State
}

func (s *GameSuite) SetUpSuite(c *C) {
	var state *State
	err := json.Unmarshal([]byte(stateStr), &state)
	if err != nil {
		panic(err)
	}
	s.state = state
	s.game = NewGame(state)
}

func (s *GameSuite) TestNewGame(c *C) {
	c.Assert(s.game.State, DeepEquals, s.state)
	c.Assert(s.state.Game, DeepEquals, s.game)
}
