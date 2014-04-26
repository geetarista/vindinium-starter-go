package vindinium

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	. "gopkg.in/check.v1"
)

var (
	_ = Suite(&ClientSuite{})
)

type ClientSuite struct {
	client *Client
	state  State
}

func (s *ClientSuite) SetUpTest(c *C) {
	s.client = &Client{}
	s.state = State{}
}

func (s *ClientSuite) TestNewClient_Training(c *C) {
	client := NewClient("server", "key", "training", "fighter", "1", false, true)
	c.Assert(client.Server, Equals, "server")
	c.Assert(client.Key, Equals, "key")
	c.Assert(client.Mode, Equals, "training")
	c.Assert(client.Bot, DeepEquals, &FighterBot{})
	c.Assert(client.Turns, Equals, "1")
	c.Assert(client.Url, Equals, "server/api/training")
}

func (s *ClientSuite) TestNewClient_Arena(c *C) {
	client := NewClient("server", "key", "arena", "random", "1", true, false)
	c.Assert(client.Server, Equals, "server")
	c.Assert(client.Key, Equals, "key")
	c.Assert(client.Mode, Equals, "arena")
	c.Assert(client.Bot, DeepEquals, &RandomBot{})
	c.Assert(client.Turns, Equals, "1")
	c.Assert(client.Url, Equals, "server/api/arena")
}

func (s *ClientSuite) TestFinished_True(c *C) {
	client := &Client{State: &State{Game: &Game{Finished: true}}}
	c.Assert(client.finished(), Equals, true)
}

func (s *ClientSuite) TestFinished_False(c *C) {
	client := &Client{State: &State{Game: &Game{Finished: false}}}
	c.Assert(client.finished(), Equals, false)
}

func (s *ClientSuite) TestMove(c *C) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"game":{"id":"derp","board":{"token":"xyz789","size":5,"tiles":"  ##[]$-@1  &&**()^^  ##[]$-@1  ##[]$-@1  ##[]$-@1"}},"playUrl":"lol"}`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	stateStr := fmt.Sprintf(`{"game":{"board":{"id":"abc123","token":"xyz789","size":5,"tiles":"  ##[]$-@1  &&**()^^  ##[]$-@1  ##[]$-@1  ##[]$-@1"}},"playUrl":"%s/api/abc123/xyz789/play"}`, server.URL)
	_ = json.Unmarshal([]byte(stateStr), &s.state)
	client := NewClient(server.URL, "key", "arena", "fighter", "1", true, false)
	client.State = &s.state
	err := client.move(Direction("East"))
	c.Assert(err, Equals, nil)
	c.Assert(client.State.Game.Id, Equals, "derp")
}

func (s *ClientSuite) TestStart_Arena(c *C) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_ = r.FormValue("key")
		_ = r.FormValue("turns")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"game":{"id":"abc123","token":"xyz789","board":{"size":2,"tiles":"  ##[]$-"}},"playUrl":"server"}`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := NewClient(server.URL, "xyz789", "arena", "fighter", "1", true, false)
	client.Start()

	c.Assert(client.Key, Equals, "xyz789")
	c.Assert(client.Mode, Equals, "arena")
	c.Assert(client.State.Game.Token, Equals, "xyz789")
}

func (s *ClientSuite) TestStart_Training(c *C) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_ = r.FormValue("key")
		_ = r.FormValue("turns")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"game":{"id":"abc123","token":"xyz789","board":{"size":2,"tiles":"  ##[]$-"}},"playUrl":"server"}`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := NewClient(server.URL, "xyz789", "training", "fighter", "10", false, true)
	client.Start()

	c.Assert(client.Key, Equals, "xyz789")
	c.Assert(client.Mode, Equals, "training")
	c.Assert(client.State.Game.Token, Equals, "xyz789")
}
func (s *ClientSuite) TestStart_BadUri(c *C) {
	handler := func(w http.ResponseWriter, r *http.Request) {}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := NewClient("nonexistent", "xyz789", "training", "fighter", "10", true, false)
	err := client.Start()

	c.Assert(err, FitsTypeOf, &url.Error{})
}

func (s *ClientSuite) TestStart_BadJson(c *C) {
	server := httptest.NewServer(http.HandlerFunc(invalidJSONHandler))
	defer server.Close()
	client := NewClient(server.URL, "xyz789", "training", "fighter", "10", true, false)
	err := client.Start()

	c.Assert(err, FitsTypeOf, &json.SyntaxError{})
}

func (s *ClientSuite) TestPlay(c *C) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		dir := r.FormValue("dir")
		finished := false
		if dir == "North" {
			finished = true
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fmt.Sprintf(`{"game":{"id":"abc123","token":"xyz789","board":{"size":2,"tiles":"  ##[]$-"},"finished":%t},"playUrl":"%s"}`, finished, s.client.State.PlayUrl))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	stateStr := fmt.Sprintf(`{"game":{"board":{"id":"abc123","token":"xyz789","size":5,"tiles":"  ##[]$-@1  &&**()^^  ##[]$-@1  ##[]$-@1  ##[]$-@1"}},"playUrl":"%s/api/abc123/xyz789/play"}`, server.URL)
	_ = json.Unmarshal([]byte(stateStr), &s.state)
	client := NewClient(server.URL, "xyz789", "training", "fighter", "10", true, true)
	s.client = client
	client.State = &s.state
	client.Play()

	c.Assert(client.State.Game.Finished, Equals, true)
}

func (s *ClientSuite) TestPlay_Error(c *C) {
	server := httptest.NewServer(http.HandlerFunc(invalidJSONHandler))
	defer server.Close()
	stateStr := fmt.Sprintf(`{"game":{"board":{"id":"abc123","token":"xyz789","size":5,"tiles":"  ##[]$-@1  &&**()^^  ##[]$-@1  ##[]$-@1  ##[]$-@1"}},"playUrl":"%s/api/abc123/xyz789/play"}`, server.URL)
	_ = json.Unmarshal([]byte(stateStr), &s.state)
	client := NewClient(server.URL, "xyz789", "training", "fighter", "10", true, false)
	s.client = client
	client.State = &s.state
	err := client.Play()

	c.Assert(err, FitsTypeOf, &json.SyntaxError{})
	c.Assert(client.State.Game.Finished, Equals, false)
}
