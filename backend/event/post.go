package event

import (
	C "github.com/carlosCACB333/cb-back/const"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/util"
	"github.com/nats-io/nats.go"
)

type PostEvent struct {
	nats *nats.Conn
}

func NewPostEvent(nats *nats.Conn) *PostEvent {
	return &PostEvent{
		nats: nats,
	}
}

func (p *PostEvent) PublishCreated(post model.Post) error {
	msg, err := util.EncodeEventMsg(post)
	if err != nil {
		return err
	}
	return p.nats.Publish(C.EVENT_POST_CREATED, msg)
}

func (p *PostEvent) OnPublishCreated(f func(post model.Post)) error {
	_, err := p.nats.Subscribe(C.EVENT_POST_CREATED, func(msg *nats.Msg) {
		var post model.Post
		err := util.DecodeEventMsg(msg.Data, &post)
		if err != nil {
			return
		}
		f(post)
	})
	return err
}
