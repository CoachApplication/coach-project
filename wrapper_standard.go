package project

import (
	"context"
	"github.com/CoachApplication/api"
)

type StandardWrapper struct {
	ctx context.Context
	ops api.Operations
}

func NewStandardWrapper(ops api.Operations, ctx context.Context) *StandardWrapper {
	if ctx == nil {
		ctx = context.Background()
	}
	return &StandardWrapper{
		ops: ops,
		ctx: ctx,
	}
}

func (sw *StandardWrapper) Wrapper() Wrapper {
	return Wrapper(sw)
}

func (sw *StandardWrapper) Name() string {
	name := ""
	if nOp, err := sw.ops.Get(OPERATION_ID_NAME); err == nil {

		res := nOp.Exec(nOp.Properties())

		select {
		case <-res.Finished():

			if res.Success() {
				if nP, err := res.Properties().Get(PROPERTY_ID_NAME); err == nil {
					if n, good := nP.Get().(string); good {
						name = n
					}
				}
			}

		case <-sw.ctx.Done():
		}
	}
	return name
}

func (sw *StandardWrapper) Label() string {
	label := ""
	if nOp, err := sw.ops.Get(OPERATION_ID_LABEL); err == nil {

		res := nOp.Exec(nOp.Properties())

		select {
		case <-res.Finished():

			if res.Success() {
				if lP, err := res.Properties().Get(PROPERTY_ID_LABEL); err == nil {
					if l, good := lP.Get().(string); good {
						label = l
					}
				}
			}

		case <-sw.ctx.Done():
		}
	}
	return label
}

func (sw *StandardWrapper) Env() map[string]string {
	env := map[string]string{}

	if nOp, err := sw.ops.Get(OPERATION_ID_ENV_MAP); err == nil {

		res := nOp.Exec(nOp.Properties())

		select {
		case <-res.Finished():

			if res.Success() {
				if eProp, err := res.Properties().Get(PROPERTY_ID_ENV_MAP); err == nil {
					if e, good := eProp.Get().(map[string]string); good {
						env = e
					}
				}
			}

		case <-sw.ctx.Done():
		}
	}

	return env
}
