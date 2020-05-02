package util

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
    a1, b1 := a.(int), b.(int)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func Int8Comparator(a, b interface{}) int {
    a1, b1 := a.(int8), b.(int8)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func Int16Comparator(a, b interface{}) int {
    a1, b1 := a.(int16), b.(int16)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func Int32Comparator(a, b interface{}) int {
    a1, b1 := a.(int32), b.(int32)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func Int64Comparator(a, b interface{}) int {
    a1, b1 := a.(int64), b.(int64)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func UIntComparator(a, b interface{}) int {
    a1, b1 := a.(uint), b.(uint)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func UInt8Comparator(a, b interface{}) int {
    a1, b1 := a.(uint8), b.(uint8)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func UInt16Comparator(a, b interface{}) int {
    a1, b1 := a.(uint16), b.(uint16)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func UInt32Comparator(a, b interface{}) int {
    a1, b1 := a.(uint32), b.(uint32)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func UInt64Comparator(a, b interface{}) int {
    a1, b1 := a.(uint64), b.(uint64)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func Float32Comparator(a, b interface{}) int {
    a1, b1 := a.(float32), b.(float32)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func Float64Comparator(a, b interface{}) int {
    a1, b1 := a.(float64), b.(float64)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func ByteComparator(a, b interface{}) int {
    a1, b1 := a.(byte), b.(byte)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}

func RuneComparator(a, b interface{}) int {
    a1, b1 := a.(rune), b.(rune)
    if a1 > b1 {
        return 1
    } else if a1 < b1 {
        return -1
    } else {
        return 0
    }
}
