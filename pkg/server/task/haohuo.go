package task

import (
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"time"
)

type Listener interface {
	OnChange(haohuos []orm.BusinessHaohuo)
}

type newHaoHuosPublisher struct {
	listeners map[Listener]bool
}

var publisher *newHaoHuosPublisher

func (p *newHaoHuosPublisher) AddListener(l Listener) {
	p.listeners[l] = true
}

func (p *newHaoHuosPublisher) RemoveListener(l Listener) {
	delete(p.listeners, l)
}

func InitTasks() {
	publisher = &newHaoHuosPublisher{
		listeners: make(map[Listener]bool),
	}

	go func() {
		for {
			time.Sleep(15 * time.Second)
			hs, e := service.FindUserBusinessHaohuosByLimit(4)
			if e != nil {
				logrus.Error(e)
				continue
			}
			for k := range publisher.listeners {
				go k.OnChange(hs)
			}
		}
	}()
}

func AddListener(l Listener) {
	publisher.AddListener(l)
}
