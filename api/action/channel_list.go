package action

import (
	"net/http"
	"github.com/Liar233/Scheduler/scheduler"
)

type ChannelListAction struct {
	cp *scheduler.ChannelPool
}

func (a *ChannelListAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	WriteSuccessResponse(w, a.cp.GetList())
}

func NewChannelListAction(cp *scheduler.ChannelPool) *ChannelListAction {
	return &ChannelListAction{
		cp: cp,
	}
}
