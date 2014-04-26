package vindinium

import (
	"fmt"
	"net/http"
	"testing"

	. "gopkg.in/check.v1"
)

var (
	invalidJSONHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{`)
	}

	stateStr = `{"game":{"id":"abc123","turn":0,"maxTurns":1200,"heroes":[{"id":1,"name":"geetarista","userId":"user123","elo":1200,"pos":{"x":0,"y":0},"life":100,"gold":0,"mineCount":0,"spawnPos":{"x":0,"y":0},"crashed":false}],"board":{"size":5,"tiles":"  ##[]$-@1  &&**()^^  ##[]$-@1  ##[]$-@1  ##[]$-@1"},"finished":false},"hero":{"id":1,"name":"geetarista","userId":"user123","elo":1200,"pos":{"x":0,"y":0},"life":100,"gold":0,"mineCount":0,"spawnPos":{"x":0,"y":0},"crashed":false},"token":"xyz789","viewUrl":"http://vindinium.org/321zyx","playUrl":"http://vindinium.org/api/321zyx/098/play"}`
)

func Test(t *testing.T) { TestingT(t) }
