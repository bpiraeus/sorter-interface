sorter-interface
================

basic sorter interface

## Sample sorter interface for map[string]<value> 

(currently int64, uint64, and float64)

let us presume you have a map[string]int64 called 'mymap'

sorted := NewInt64Sorter(mymap)
sorted.Sort()

Sort (by design) returns in descending order instead of ascending like the default, to change this
behavior you would merely flip the > to a < in the associated Less() function

returns an indexed slice of the following struct

type <ValueType>Sort struct {
  Keys []string
  Vals []ValueType
}

This could probably be made more generic using reflection but I chose to be explicit
