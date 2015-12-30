package react

import "fmt"

const TestVersion = 3

type sheet []cell
type callback *func(int)

type cell struct {
    updateList []callback
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
    for _, updateID := range c.updateList {
        update := *updateID
        update(val)
    }
}

func (c *cell) AddCallback(f func(int)) CallbackHandle {
    c.updateList = append(c.updateList, &f)
    return &f
}

func (c *cell) RemoveCallback(remove CallbackHandle) {
    var newList []callback
    for _, handle := range c.updateList {
        if handle != remove.(callback) {
            fmt.Printf("Checking %s != %s\n", handle, remove)
            newList = append(newList, handle)
        }
    }
    c.updateList = newList
}

func New() Reactor {
    return &sheet{}
}

func (s *sheet) CreateInput(val int) InputCell {
    c := cell{}
    c.SetValue(val)
    *s = append(*s, c)
    return &c
}

func (s *sheet) CreateCompute1(c Cell, f func(int)int) ComputeCell {
    inputCell := c.(*cell)
    computeCell := cell{}
    var update = func(n int) {
        computeCell.SetValue(f(n))
    }
    update(inputCell.Value())
    inputCell.updateList = append(inputCell.updateList, &update)
    return &computeCell
}

func (s *sheet) CreateCompute2(c1, c2 Cell, f func(int, int)int) ComputeCell {
    inputCell1 := c1.(*cell)
    inputCell2 := c2.(*cell)
    computeCell := cell{}
    var update1 = func(n int) {
        computeCell.SetValue(f(n, inputCell2.Value()))
    }
    update1(inputCell1.Value())
    inputCell1.updateList = append(inputCell1.updateList, &update1)

    var update2 = func(n int) {
        computeCell.SetValue(f(inputCell1.Value(), n))
    }
    update2(inputCell2.Value())
    inputCell2.updateList = append(inputCell2.updateList, &update2)

    return &computeCell
}
