package react

//TestVersion is the unit tests this will pass.
const TestVersion = 3

//sheet is a collection of cells.
type sheet struct {
	cells []*cell
}

//callback is a function that is called when the cell changes.
type callback *func(int)

//cell holds a value either computed or static.
type cell struct {
	sheet     *sheet
	value     int
	update    func() bool
	stale     bool
	callbacks []callback
}

//Value returns the value of the cell
func (c *cell) Value() int {
	return c.value
}

//SetValue sets the static value of a cell.
func (c *cell) SetValue(val int) {
	c.stale = (c.Value() == val)
	c.value = val
	c.sheet.update()
}

/*AddCallback adds a function to be run when the cell changes*/
func (c *cell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks = append(c.callbacks, &f)
	return &f
}

/*RemoveCallback filters out callbacks so that the function no longer runs
when a cell changes.*/
func (c *cell) RemoveCallback(remove CallbackHandle) {
	var newCallbacks []callback
	for _, handle := range c.callbacks {
		if handle != callback(remove.(*func(int))) {
			newCallbacks = append(newCallbacks, handle)
		}
	}
	c.callbacks = newCallbacks
}

/*New creats a new collection of cells.*/
func New() Reactor {
	return &sheet{}
}

/*CreateInput adds an cell with a static value to the collection of cells.*/
func (s *sheet) CreateInput(val int) InputCell {
	inpCell := cell{sheet: s, value: val}

	inpCell.update = func() bool {
		state := inpCell.stale
		inpCell.stale = false
		return state
	}

	s.cells = append(s.cells, &inpCell)
	return &inpCell
}

/*CreateCompute1 adds a cell with a computed value to the collection of cells.*/
func (s *sheet) CreateCompute1(input Cell, compVal func(int) int) ComputeCell {
	var comp = func(cells []Cell) int {
		return compVal(cells[0].Value())
	}

	return s.createComputeGeneral([]Cell{input}, comp)
}

/*CreateCompute2 adds a cell with a computed value that depends on two cells.*/
func (s *sheet) CreateCompute2(input1, input2 Cell, compVal func(int, int) int) ComputeCell {
	var comp = func(cells []Cell) int {
		return compVal(cells[0].Value(), cells[1].Value())
	}

	return s.createComputeGeneral([]Cell{input1, input2}, comp)
}

/*createComputeGeneral creates a compute cell that can depend on
any number of cells.*/
func (s *sheet) createComputeGeneral(cells []Cell, compFunc func([]Cell) int) ComputeCell {
	// Like an input cell but with a differnt update function
	compCell := s.CreateInput(0).(*cell)

	compCell.update = func() bool {
		oldVal := compCell.Value()
		compCell.value = compFunc(cells)
		return oldVal != compCell.Value()
	}

	compCell.update()
	return compCell
}

/*update checks all the cells in the sheet for an changes and runs callback on all
changed cells*/
func (s *sheet) update() {
	for _, cellID := range s.cells {
		if cellID.update() {
			for _, callID := range cellID.callbacks {
				call := *callID
				call(cellID.Value())
			}
		}
	}
}
