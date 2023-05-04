package state

import "fmt"

type Context struct {
	state State
}

func (c *Context) ChangeState(state State) {
	c.state = state
}

func (c *Context) Start() {
	c.state.Start(c)
}

func (c *Context) Stop() {
	c.state.Stop(c)
}

func (c *Context) Run() {
	c.state.Run(c)
}

type State interface {
	Start(ctx *Context)
	Stop(ctx *Context)
	Run(ctx *Context)
}

type StartState struct {
}

func (s *StartState) Start(ctx *Context) {
	fmt.Println("already start")
}

func (s *StartState) Stop(ctx *Context) {
	fmt.Println("change start state to stop")
	ctx.ChangeState(&StopState{})
}

func (s *StartState) Run(ctx *Context) {
	fmt.Println("change start state to run")
	ctx.ChangeState(&RunState{})
}

type StopState struct {
}

func (s *StopState) Start(ctx *Context) {
	fmt.Println("change stop state to start")
	ctx.ChangeState(&StartState{})
}

func (s *StopState) Stop(ctx *Context) {
	fmt.Println("already stop")
}

func (s *StopState) Run(ctx *Context) {
	fmt.Println("can't change stop state to run")
}

type RunState struct {
}

func (s *RunState) Start(ctx *Context) {
	fmt.Println("can't change run state to start")
}

func (s *RunState) Stop(ctx *Context) {
	fmt.Println("change run state to stop")
	ctx.ChangeState(&StopState{})
}

func (s *RunState) Run(ctx *Context) {
	fmt.Println("already run")
}
