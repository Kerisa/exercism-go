package listops

// IntList doc here...
type IntList    []int
type binFunc    func(x, y int)  int
type predFunc   func(x int)     bool
type unaryFunc  func(x int)     int


// Foldl doc...
func (L *IntList) Foldl(fn func(x, y int) int, initial int) int {
    result := initial
    for _,tt := range(*L) {
        result = fn(result, tt)
    }
    return result
}

// Foldr doc...
func (L *IntList) Foldr(fn func(x, y int) int, initial int) int {
    result := initial
    for i := len(*L) - 1; i >= 0; i-- {
        result = fn((*L)[i], result)
    }
    return result
}

// Filter doc...
func (L *IntList) Filter(fn func(x int) bool) IntList {
    newIntList := IntList{}
    for _,tt := range(*L) {
        if fn(tt) {
            newIntList = append(newIntList, tt)
        }
    }
    return newIntList
}

// Length doc...
func (L *IntList) Length() int {
    return len(*L)
}

// Map doc...
func (L *IntList) Map(fn func(x int) int) IntList {
    result := IntList{}
    for _,tt := range(*L) {
        result = append(result, fn(tt))
    }
    return result
}

// Reverse doc...
func (L *IntList) Reverse() IntList {
    result := IntList{}
    for i := len(*L) - 1; i >= 0; i-- {
        result = append(result, (*L)[i])
    }
    return result
}

// Append doc...
func (L *IntList) Append(newPart IntList) IntList {
    result := *L
    for _,tt := range(newPart) {
        result = append(result, tt)
    }
    return result
}

// Concat doc...
func (L *IntList) Concat(newPart []IntList) IntList {
    result := *L
    for _,tt := range(newPart) {
        result = result.Append(tt)
    }
    return result
}
