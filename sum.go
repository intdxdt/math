package math

//Sum float64s
func SumF64(o []float64) float64 {
    var t float64
    for _, v := range o {
        t += v
    }
    return t
}

//Sum int64s
func SumI64(o []int64) int64 {
    var t int64
    for _, v := range o {
        t += v
    }
    return t
}

//Sum ints
func SumInt(o []int) int {
    var t int
    for _, v := range o {
        t += v
    }
    return t
}
