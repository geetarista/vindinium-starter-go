package vindinium

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

var (
	_ = Suite(&BoardSuite{})
)

type BoardSuite struct {
	board *Board
	state *State
}

func (s *BoardSuite) SetUpSuite(c *C) {
	var state *State
	err := json.Unmarshal([]byte(stateStr), &state)
	if err != nil {
		panic(err)
	}
	s.state = state
	s.board = NewGame(state).Board
}

func (s *BoardSuite) TestBadTile(c *C) {
	c.Assert(s.board.Tileset[1][2], Equals, -3)
}

func (s *BoardSuite) TestParse(c *C) {
	c.Assert(s.board.Tileset[0][0], Equals, AIR)
	c.Assert(s.board.Tileset[0][1], Equals, WALL)
	c.Assert(s.board.Tileset[0][2], Equals, TAVERN)
	gotMine := s.board.Tileset[0][3]
	wantMine := &MineTile{string([]rune(s.board.Tiles)[7])}
	c.Assert(gotMine, FitsTypeOf, &MineTile{})
	c.Assert(gotMine.(*MineTile).HeroId, Equals, wantMine.HeroId)
	id := tileToInt(s.board.Tiles, 9)
	c.Assert(s.board.Tileset[0][4].(*HeroTile).Id, DeepEquals, (&HeroTile{id}).Id)
}

func (s *BoardSuite) TestPassable(c *C) {
	c.Assert(s.board.Passable(Position{0, 0}), Equals, true)  // air
	c.Assert(s.board.Passable(Position{0, 1}), Equals, false) // wall
	c.Assert(s.board.Passable(Position{0, 2}), Equals, false) // tavern
	c.Assert(s.board.Passable(Position{0, 3}), Equals, true)  // mine
	c.Assert(s.board.Passable(Position{0, 4}), Equals, true)  // hero
}

func (s *BoardSuite) TestTo(c *C) {
	c.Assert(s.board.To(Position{0, 0}, "North"), DeepEquals, &Position{0, 0})
	c.Assert(s.board.To(Position{1, 1}, "North"), DeepEquals, &Position{0, 1})
	c.Assert(s.board.To(Position{2, 2}, "North"), DeepEquals, &Position{1, 2})
	c.Assert(s.board.To(Position{6, 0}, "South"), DeepEquals, &Position{5, 0})
	c.Assert(s.board.To(Position{0, 0}, "South"), DeepEquals, &Position{1, 0})
	c.Assert(s.board.To(Position{1, 2}, "South"), DeepEquals, &Position{2, 2})
	c.Assert(s.board.To(Position{0, 6}, "East"), DeepEquals, &Position{0, 5})
	c.Assert(s.board.To(Position{0, 0}, "East"), DeepEquals, &Position{0, 1})
	c.Assert(s.board.To(Position{1, 1}, "East"), DeepEquals, &Position{1, 2})
	c.Assert(s.board.To(Position{2, 0}, "West"), DeepEquals, &Position{2, 0})
	c.Assert(s.board.To(Position{1, 1}, "West"), DeepEquals, &Position{1, 0})
	c.Assert(s.board.To(Position{1, 2}, "West"), DeepEquals, &Position{1, 1})
}
