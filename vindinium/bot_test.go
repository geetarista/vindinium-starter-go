package vindinium

import (
	"encoding/json"
	. "gopkg.in/check.v1"
)

var (
	_ = Suite(&BotSuite{})
)

type BotSuite struct {
	state *State
}

func (s *BotSuite) SetUpSuite(c *C) {
	var state *State
	err := json.Unmarshal([]byte(stateStr), &state)
	if err != nil {
		panic(err)
	}
	s.state = state
}

func (s *BotSuite) TestRandomBot(c *C) {
	bot := &RandomBot{}
	dir := bot.Move(s.state)
	c.Assert(dir, FitsTypeOf, Direction("East"))
}

func (s *BotSuite) TestFighterBot(c *C) {
	bot := &FighterBot{}
	dir := bot.Move(s.state)
	c.Assert(dir, FitsTypeOf, Direction("East"))
}
