// moving sorting tools into their own file

package sorter

// import necessary packages
import (
	"sort"
)

// int64 sort
type Uint64Sort struct {
	Keys []string
	Vals []uint64
}

func NewUint64Sorter(m map[string]uint64) *Uint64Sort {

	vs := &Uint64Sort{
		Keys: make([]string, 0, len(m)),
		Vals: make([]uint64, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *Uint64Sort) Sort() {
	sort.Sort(vs)
}

func (vs *Uint64Sort) Len() int           { return len(vs.Vals) }
func (vs *Uint64Sort) Less(i, j int) bool { return vs.Vals[i] > vs.Vals[j] }

func (vs *Uint64Sort) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}

// float64 sort
type Float64Sort struct {
	Keys []string
	Vals []float64
}

func NewFloat64Sorter(m map[string]float64) *Float64Sort {

	vs := &Float64Sort{
		Keys: make([]string, 0, len(m)),
		Vals: make([]float64, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *Float64Sort) Sort() {
	sort.Sort(vs)
}

func (vs *Float64Sort) Len() int           { return len(vs.Vals) }
func (vs *Float64Sort) Less(i, j int) bool { return vs.Vals[i] > vs.Vals[j] }

func (vs *Float64Sort) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}

// int64 sort
type Int64Sort struct {
	Keys []string
	Vals []int64
}

func NewInt64Sorter(m map[string]int64) *Int64Sort {

	vs := &Int64Sort{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int64, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *Int64Sort) Sort() {
	sort.Sort(vs)
}

func (vs *Int64Sort) Len() int           { return len(vs.Vals) }
func (vs *Int64Sort) Less(i, j int) bool { return vs.Vals[i] > vs.Vals[j] }

func (vs *Int64Sort) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}
