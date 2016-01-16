package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

//Matrix stores an array of numbers in rows and columns
type Matrix struct {
	rows, cols int
	data       []int
}

/*New creates a matrix from a string.*/
func New(str string) (*Matrix, error) {
	rows := strings.Split(str, "\n")
	matrix := Matrix{rows: len(rows), cols: -1}
	for _, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		if matrix.cols == -1 {
			matrix.cols = len(cols)
		} else if matrix.cols != len(cols) {
			return nil, fmt.Errorf("Rows need to be the same length")
		}
		for _, char := range cols {
			num, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			matrix.data = append(matrix.data, num)
		}
	}
	return &matrix, nil
}

/*Rows gets the matrix represented in rows.*/
func (m Matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for r := 0; r < m.rows; r++ {
		rows[r] = make([]int, m.cols)
		for c := 0; c < m.cols; c++ {
			rows[r][c] = m.data[r*m.cols+c]
		}
	}
	return rows
}

/*Cols gets the matrix represented in columns*/
func (m Matrix) Cols() [][]int {
	cols := make([][]int, m.cols)
	for c := 0; c < m.cols; c++ {
		cols[c] = make([]int, m.rows)
		for r := 0; r < m.rows; r++ {
			cols[c][r] = m.data[r*m.cols+c]
		}
	}
	return cols
}

/*Set sets the value of the matrix at point row, col.*/
func (m *Matrix) Set(row, col, value int) bool {
	if row < 0 || m.rows <= row || col < 0 || m.cols <= col {
		return false
	}
	m.data[row*m.cols+col] = value
	return true
}
