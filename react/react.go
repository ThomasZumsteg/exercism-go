package react

// import "fmt"

const TestVersion = 3

type sheet []cell
type callback *func(int)

type cell struct {
    callbacks []callback
    value int
}

func (c *cell) Value() int {
    return c.value
}

func (c *cell) SetValue(val int) {
    if c.value == val {
        return
    }
    c.value = val
    for _, callID := range c.callbacks{
        call := *callID
        call(val)
    }
}

func (c *cell) AddCallback(call func(int)) CallbackHandle {
    c.callbacks = append(c.callbacks, &call)
    return &call
}

func (c *cell) RemoveCallback(remove CallbackHandle) {
    var newCallBacks []callback
    removeAddr := callback(remove.(*func(int)))
    for _, handle := range c.callbacks {
        if handle != removeAddr {
            newCallBacks = append(newCallBacks, handle)
        }
    }
    c.callbacks = newCallBacks
}

func New() Reactor {
    return &sheet{}
}

func (s *sheet) CreateInput(val int) InputCell {
    inputCell := cell{}
    inputCell.SetValue(val)
    *s = append(*s, inputCell)
    return &inputCell
}

func (s *sheet) CreateCompute1(c Cell, f func(int)int) ComputeCell {
    inputCell := c.(*cell)
    computeCell := cell{}
    var update = func(_ int) {
        computeCell.SetValue(f(inputCell.Value()))
    }
    inputCell.AddCallback( update )
    update(0)
    return &computeCell
}

func (s *sheet) CreateCompute2(c1, c2 Cell, f func(int, int)int) ComputeCell {
    inputCell1 := c1.(*cell)
    inputCell2 := c2.(*cell)
    computeCell := cell{}

    var update = func(_ int) {
        computeCell.SetValue(f(inputCell1.Value(), inputCell2.Value()))
    }
    update(0)
    inputCell1.AddCallback( update )
    inputCell2.AddCallback( update )

    return &computeCell
}
