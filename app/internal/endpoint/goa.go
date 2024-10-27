package endpoint

import (
	"goa.design/clue/debug"
	"goa.design/clue/log"
	goa "goa.design/goa/v3/pkg"

	"github.com/jace-ys/countup/internal/endpoint/middleware/goaerror"
	"github.com/jace-ys/countup/internal/endpoint/middleware/tracer"
)

type GoaAdapter[S any, E GoaEndpoints] struct {
	newFunc GoaNewFunc[S, E]
}

type GoaEndpoints interface {
	Use(func(goa.Endpoint) goa.Endpoint)
}

type GoaNewFunc[S any, E GoaEndpoints] func(svc S) E

func Goa[S any, E GoaEndpoints](newFunc GoaNewFunc[S, E]) *GoaAdapter[S, E] {
	return &GoaAdapter[S, E]{
		newFunc: newFunc,
	}
}

func (a *GoaAdapter[S, E]) Adapt(svc S) E {
	ep := a.newFunc(svc)

	chainMiddleware(ep,
		tracer.Endpoint,
		log.Endpoint,
		debug.LogPayloads(),
		goaerror.Reporter,
	)

	return ep
}

func chainMiddleware[E GoaEndpoints](ep E, m ...func(goa.Endpoint) goa.Endpoint) {
	for i := len(m) - 1; i >= 0; i-- {
		ep.Use(m[i])
	}
}
