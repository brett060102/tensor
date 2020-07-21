// Code generated by genlib2. DO NOT EDIT.

package native

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
	. "gorgonia.org/tensor"
)

func checkNativeIterable(t *Dense, dims int, dt Dtype) error {
	// checks:
	if !t.IsNativelyAccessible() {
		return errors.Errorf("Cannot convert *Dense to *mat.Dense. Data is inaccessible")
	}

	if t.Shape().Dims() != dims {
		return errors.Errorf("Cannot convert *Dense to native iterator. Expected number of dimension: %d, T has got %d dimensions (Shape: %v)", dims, t.Dims(), t.Shape())
	}

	if t.F() || t.RequiresIterator() {
		return errors.Errorf("Not yet implemented: native matrix for colmajor or unpacked matrices")
	}

	if t.Dtype() != dt {
		return errors.Errorf("Conversion to native iterable only works on %v. Got %v", dt, t.Dtype())
	}

	return nil
}

/* Native Iterables for bool */

// VectorB converts a *Dense into a []bool
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorB(t *Dense) (retVal []bool, err error) {
	if err = checkNativeIterable(t, 1, Bool); err != nil {
		return nil, err
	}
	return t.Bools(), nil
}

// MatrixB converts a  *Dense into a [][]bool
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixB(t *Dense) (retVal [][]bool, err error) {
	if err = checkNativeIterable(t, 2, Bool); err != nil {
		return nil, err
	}

	data := t.Bools()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]bool, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]bool, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3B converts a *Dense into a [][][]bool.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3B(t *Dense) (retVal [][][]bool, err error) {
	if err = checkNativeIterable(t, 3, Bool); err != nil {
		return nil, err
	}

	data := t.Bools()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]bool, layers)
	for i := range retVal {
		retVal[i] = make([][]bool, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]bool, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for int */

// VectorI converts a *Dense into a []int
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorI(t *Dense) (retVal []int, err error) {
	if err = checkNativeIterable(t, 1, Int); err != nil {
		return nil, err
	}
	return t.Ints(), nil
}

// MatrixI converts a  *Dense into a [][]int
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixI(t *Dense) (retVal [][]int, err error) {
	if err = checkNativeIterable(t, 2, Int); err != nil {
		return nil, err
	}

	data := t.Ints()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]int, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]int, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3I converts a *Dense into a [][][]int.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3I(t *Dense) (retVal [][][]int, err error) {
	if err = checkNativeIterable(t, 3, Int); err != nil {
		return nil, err
	}

	data := t.Ints()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]int, layers)
	for i := range retVal {
		retVal[i] = make([][]int, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]int, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for int8 */

// VectorI8 converts a *Dense into a []int8
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorI8(t *Dense) (retVal []int8, err error) {
	if err = checkNativeIterable(t, 1, Int8); err != nil {
		return nil, err
	}
	return t.Int8s(), nil
}

// MatrixI8 converts a  *Dense into a [][]int8
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixI8(t *Dense) (retVal [][]int8, err error) {
	if err = checkNativeIterable(t, 2, Int8); err != nil {
		return nil, err
	}

	data := t.Int8s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]int8, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]int8, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3I8 converts a *Dense into a [][][]int8.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3I8(t *Dense) (retVal [][][]int8, err error) {
	if err = checkNativeIterable(t, 3, Int8); err != nil {
		return nil, err
	}

	data := t.Int8s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]int8, layers)
	for i := range retVal {
		retVal[i] = make([][]int8, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]int8, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for int16 */

// VectorI16 converts a *Dense into a []int16
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorI16(t *Dense) (retVal []int16, err error) {
	if err = checkNativeIterable(t, 1, Int16); err != nil {
		return nil, err
	}
	return t.Int16s(), nil
}

// MatrixI16 converts a  *Dense into a [][]int16
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixI16(t *Dense) (retVal [][]int16, err error) {
	if err = checkNativeIterable(t, 2, Int16); err != nil {
		return nil, err
	}

	data := t.Int16s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]int16, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]int16, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3I16 converts a *Dense into a [][][]int16.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3I16(t *Dense) (retVal [][][]int16, err error) {
	if err = checkNativeIterable(t, 3, Int16); err != nil {
		return nil, err
	}

	data := t.Int16s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]int16, layers)
	for i := range retVal {
		retVal[i] = make([][]int16, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]int16, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for int32 */

// VectorI32 converts a *Dense into a []int32
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorI32(t *Dense) (retVal []int32, err error) {
	if err = checkNativeIterable(t, 1, Int32); err != nil {
		return nil, err
	}
	return t.Int32s(), nil
}

// MatrixI32 converts a  *Dense into a [][]int32
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixI32(t *Dense) (retVal [][]int32, err error) {
	if err = checkNativeIterable(t, 2, Int32); err != nil {
		return nil, err
	}

	data := t.Int32s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]int32, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]int32, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3I32 converts a *Dense into a [][][]int32.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3I32(t *Dense) (retVal [][][]int32, err error) {
	if err = checkNativeIterable(t, 3, Int32); err != nil {
		return nil, err
	}

	data := t.Int32s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]int32, layers)
	for i := range retVal {
		retVal[i] = make([][]int32, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]int32, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for int64 */

// VectorI64 converts a *Dense into a []int64
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorI64(t *Dense) (retVal []int64, err error) {
	if err = checkNativeIterable(t, 1, Int64); err != nil {
		return nil, err
	}
	return t.Int64s(), nil
}

// MatrixI64 converts a  *Dense into a [][]int64
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixI64(t *Dense) (retVal [][]int64, err error) {
	if err = checkNativeIterable(t, 2, Int64); err != nil {
		return nil, err
	}

	data := t.Int64s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]int64, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]int64, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3I64 converts a *Dense into a [][][]int64.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3I64(t *Dense) (retVal [][][]int64, err error) {
	if err = checkNativeIterable(t, 3, Int64); err != nil {
		return nil, err
	}

	data := t.Int64s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]int64, layers)
	for i := range retVal {
		retVal[i] = make([][]int64, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]int64, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for uint */

// VectorU converts a *Dense into a []uint
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorU(t *Dense) (retVal []uint, err error) {
	if err = checkNativeIterable(t, 1, Uint); err != nil {
		return nil, err
	}
	return t.Uints(), nil
}

// MatrixU converts a  *Dense into a [][]uint
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixU(t *Dense) (retVal [][]uint, err error) {
	if err = checkNativeIterable(t, 2, Uint); err != nil {
		return nil, err
	}

	data := t.Uints()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]uint, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]uint, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3U converts a *Dense into a [][][]uint.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3U(t *Dense) (retVal [][][]uint, err error) {
	if err = checkNativeIterable(t, 3, Uint); err != nil {
		return nil, err
	}

	data := t.Uints()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]uint, layers)
	for i := range retVal {
		retVal[i] = make([][]uint, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]uint, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for uint8 */

// VectorU8 converts a *Dense into a []uint8
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorU8(t *Dense) (retVal []uint8, err error) {
	if err = checkNativeIterable(t, 1, Uint8); err != nil {
		return nil, err
	}
	return t.Uint8s(), nil
}

// MatrixU8 converts a  *Dense into a [][]uint8
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixU8(t *Dense) (retVal [][]uint8, err error) {
	if err = checkNativeIterable(t, 2, Uint8); err != nil {
		return nil, err
	}

	data := t.Uint8s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]uint8, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]uint8, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3U8 converts a *Dense into a [][][]uint8.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3U8(t *Dense) (retVal [][][]uint8, err error) {
	if err = checkNativeIterable(t, 3, Uint8); err != nil {
		return nil, err
	}

	data := t.Uint8s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]uint8, layers)
	for i := range retVal {
		retVal[i] = make([][]uint8, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]uint8, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for uint16 */

// VectorU16 converts a *Dense into a []uint16
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorU16(t *Dense) (retVal []uint16, err error) {
	if err = checkNativeIterable(t, 1, Uint16); err != nil {
		return nil, err
	}
	return t.Uint16s(), nil
}

// MatrixU16 converts a  *Dense into a [][]uint16
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixU16(t *Dense) (retVal [][]uint16, err error) {
	if err = checkNativeIterable(t, 2, Uint16); err != nil {
		return nil, err
	}

	data := t.Uint16s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]uint16, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]uint16, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3U16 converts a *Dense into a [][][]uint16.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3U16(t *Dense) (retVal [][][]uint16, err error) {
	if err = checkNativeIterable(t, 3, Uint16); err != nil {
		return nil, err
	}

	data := t.Uint16s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]uint16, layers)
	for i := range retVal {
		retVal[i] = make([][]uint16, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]uint16, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for uint32 */

// VectorU32 converts a *Dense into a []uint32
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorU32(t *Dense) (retVal []uint32, err error) {
	if err = checkNativeIterable(t, 1, Uint32); err != nil {
		return nil, err
	}
	return t.Uint32s(), nil
}

// MatrixU32 converts a  *Dense into a [][]uint32
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixU32(t *Dense) (retVal [][]uint32, err error) {
	if err = checkNativeIterable(t, 2, Uint32); err != nil {
		return nil, err
	}

	data := t.Uint32s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]uint32, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]uint32, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3U32 converts a *Dense into a [][][]uint32.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3U32(t *Dense) (retVal [][][]uint32, err error) {
	if err = checkNativeIterable(t, 3, Uint32); err != nil {
		return nil, err
	}

	data := t.Uint32s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]uint32, layers)
	for i := range retVal {
		retVal[i] = make([][]uint32, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]uint32, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for uint64 */

// VectorU64 converts a *Dense into a []uint64
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorU64(t *Dense) (retVal []uint64, err error) {
	if err = checkNativeIterable(t, 1, Uint64); err != nil {
		return nil, err
	}
	return t.Uint64s(), nil
}

// MatrixU64 converts a  *Dense into a [][]uint64
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixU64(t *Dense) (retVal [][]uint64, err error) {
	if err = checkNativeIterable(t, 2, Uint64); err != nil {
		return nil, err
	}

	data := t.Uint64s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]uint64, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]uint64, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3U64 converts a *Dense into a [][][]uint64.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3U64(t *Dense) (retVal [][][]uint64, err error) {
	if err = checkNativeIterable(t, 3, Uint64); err != nil {
		return nil, err
	}

	data := t.Uint64s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]uint64, layers)
	for i := range retVal {
		retVal[i] = make([][]uint64, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]uint64, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for float32 */

// VectorF32 converts a *Dense into a []float32
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorF32(t *Dense) (retVal []float32, err error) {
	if err = checkNativeIterable(t, 1, Float32); err != nil {
		return nil, err
	}
	return t.Float32s(), nil
}

// MatrixF32 converts a  *Dense into a [][]float32
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixF32(t *Dense) (retVal [][]float32, err error) {
	if err = checkNativeIterable(t, 2, Float32); err != nil {
		return nil, err
	}

	data := t.Float32s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]float32, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]float32, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3F32 converts a *Dense into a [][][]float32.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3F32(t *Dense) (retVal [][][]float32, err error) {
	if err = checkNativeIterable(t, 3, Float32); err != nil {
		return nil, err
	}

	data := t.Float32s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]float32, layers)
	for i := range retVal {
		retVal[i] = make([][]float32, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]float32, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for float64 */

// VectorF64 converts a *Dense into a []float64
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorF64(t *Dense) (retVal []float64, err error) {
	if err = checkNativeIterable(t, 1, Float64); err != nil {
		return nil, err
	}
	return t.Float64s(), nil
}

// MatrixF64 converts a  *Dense into a [][]float64
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixF64(t *Dense) (retVal [][]float64, err error) {
	if err = checkNativeIterable(t, 2, Float64); err != nil {
		return nil, err
	}

	data := t.Float64s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]float64, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]float64, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3F64 converts a *Dense into a [][][]float64.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3F64(t *Dense) (retVal [][][]float64, err error) {
	if err = checkNativeIterable(t, 3, Float64); err != nil {
		return nil, err
	}

	data := t.Float64s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]float64, layers)
	for i := range retVal {
		retVal[i] = make([][]float64, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]float64, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for complex64 */

// VectorC64 converts a *Dense into a []complex64
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorC64(t *Dense) (retVal []complex64, err error) {
	if err = checkNativeIterable(t, 1, Complex64); err != nil {
		return nil, err
	}
	return t.Complex64s(), nil
}

// MatrixC64 converts a  *Dense into a [][]complex64
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixC64(t *Dense) (retVal [][]complex64, err error) {
	if err = checkNativeIterable(t, 2, Complex64); err != nil {
		return nil, err
	}

	data := t.Complex64s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]complex64, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]complex64, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3C64 converts a *Dense into a [][][]complex64.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3C64(t *Dense) (retVal [][][]complex64, err error) {
	if err = checkNativeIterable(t, 3, Complex64); err != nil {
		return nil, err
	}

	data := t.Complex64s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]complex64, layers)
	for i := range retVal {
		retVal[i] = make([][]complex64, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]complex64, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for complex128 */

// VectorC128 converts a *Dense into a []complex128
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorC128(t *Dense) (retVal []complex128, err error) {
	if err = checkNativeIterable(t, 1, Complex128); err != nil {
		return nil, err
	}
	return t.Complex128s(), nil
}

// MatrixC128 converts a  *Dense into a [][]complex128
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixC128(t *Dense) (retVal [][]complex128, err error) {
	if err = checkNativeIterable(t, 2, Complex128); err != nil {
		return nil, err
	}

	data := t.Complex128s()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]complex128, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]complex128, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3C128 converts a *Dense into a [][][]complex128.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3C128(t *Dense) (retVal [][][]complex128, err error) {
	if err = checkNativeIterable(t, 3, Complex128); err != nil {
		return nil, err
	}

	data := t.Complex128s()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]complex128, layers)
	for i := range retVal {
		retVal[i] = make([][]complex128, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]complex128, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}

/* Native Iterables for string */

// VectorStr converts a *Dense into a []string
// If the *Dense does not represent a vector of the wanted type, it will return
// an error.
func VectorStr(t *Dense) (retVal []string, err error) {
	if err = checkNativeIterable(t, 1, String); err != nil {
		return nil, err
	}
	return t.Strings(), nil
}

// MatrixStr converts a  *Dense into a [][]string
// If the *Dense does not represent a matrix of the wanted type, it
// will return an error.
func MatrixStr(t *Dense) (retVal [][]string, err error) {
	if err = checkNativeIterable(t, 2, String); err != nil {
		return nil, err
	}

	data := t.Strings()
	shape := t.Shape()
	strides := t.Strides()

	rows := shape[0]
	cols := shape[1]
	rowStride := strides[0]
	retVal = make([][]string, rows)
	for i := range retVal {
		start := i * rowStride
		retVal[i] = make([]string, 0)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i]))
		hdr.Data = uintptr(unsafe.Pointer(&data[start]))
		hdr.Cap = cols
		hdr.Len = cols
	}
	return
}

// Tensor3Str converts a *Dense into a [][][]string.
// If the *Dense does not represent a 3-tensor of the wanted type, it will return an error.
func Tensor3Str(t *Dense) (retVal [][][]string, err error) {
	if err = checkNativeIterable(t, 3, String); err != nil {
		return nil, err
	}

	data := t.Strings()
	shape := t.Shape()
	strides := t.Strides()

	layers := shape[0]
	rows := shape[1]
	cols := shape[2]
	layerStride := strides[0]
	rowStride := strides[1]
	retVal = make([][][]string, layers)
	for i := range retVal {
		retVal[i] = make([][]string, rows)
		for j := range retVal[i] {
			retVal[i][j] = make([]string, 0)
			start := i*layerStride + j*rowStride
			hdr := (*reflect.SliceHeader)(unsafe.Pointer(&retVal[i][j]))
			hdr.Data = uintptr(unsafe.Pointer(&data[start]))
			hdr.Cap = cols
			hdr.Len = cols
		}
	}
	return
}
