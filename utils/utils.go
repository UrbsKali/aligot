package utils

import "fmt"

/* TODO
- [ ] Add MatrixAdd() function
- [ ] Add MatrixSubtract() function
- [ ] Add MatrixDivide() function
- [ ] Add MatrixScalarMultiply() function
- [ ] Add MatrixScalarDivide() function
- [ ] Add MatrixDotProduct() function
- [ ] Add MatrixCrossProduct() function
- [ ] Add MatrixDeterminant() function
- [ ] Add MatrixInverse() function
- [ ] Add MatrixIdentity() function
- [ ] Add MatrixPower() function
- [ ] Add MatrixTrace() function

- [ ] Convert to GPU operations
*/

type Matrix [][]float64

func (m Matrix) String() string {
	s := ""
	for _, row := range m {
		s += fmt.Sprintf("%v\n", row)
	}
	return s
}

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	return len(m[0])
}

func MatrixEqual(m, n Matrix) bool {
	if m.Rows() != n.Rows() || m.Cols() != n.Cols() {
		return false
	}
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Transpose() Matrix {
	result := make(Matrix, m.Cols())
	for i := 0; i < m.Cols(); i++ {
		result[i] = make([]float64, m.Rows())
		for j := 0; j < m.Rows(); j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func (m *Matrix) Multiply(n *Matrix) Matrix {
	result := make(Matrix, m.Rows())
	// REMOVE FOR FUTURE USE | MEMORY EFFICIENCY --
	m_ := *m
	n_ := *n
	// ------------------------------------------
	for i := 0; i < m.Rows(); i++ {
		result[i] = make([]float64, n.Cols())
		for j := 0; j < n.Cols(); j++ {
			for k := 0; k < m.Cols(); k++ {
				result[i][j] += m_[i][k] * n_[k][j]
			}
		}
	}
	return result
}
