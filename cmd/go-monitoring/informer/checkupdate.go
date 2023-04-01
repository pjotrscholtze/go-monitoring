package informer

import (
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
)

type checkUpdateInformer struct {
	listeners []func(result entity.Result, target config.Target, check config.Check)
}

type CheckUpdateListener interface {
	InformAboutCheckUpdate(result entity.Result, target config.Target, check config.Check)
}

type CheckUpdateInformer interface {
	Inform(result entity.Result, target config.Target, check config.Check)
	RegisterListenerInterface(listener CheckUpdateListener)
	RegisterListenerFunc(listener func(result entity.Result, target config.Target, check config.Check))
}

func (cui checkUpdateInformer) Inform(result entity.Result, target config.Target, check config.Check) {
	for _, listener := range cui.listeners {
		listener(result, target, check)
	}
}

func (cui *checkUpdateInformer) RegisterListenerInterface(listener CheckUpdateListener) {
	cui.listeners = append(cui.listeners, listener.InformAboutCheckUpdate)
}
func (cui *checkUpdateInformer) RegisterListenerFunc(listener func(result entity.Result, target config.Target, check config.Check)) {
	cui.listeners = append(cui.listeners, listener)
}

func NewCheckUpdateInformer() CheckUpdateInformer {
	return &checkUpdateInformer{
		listeners: []func(result entity.Result, target config.Target, check config.Check){},
	}
}
