//并行赋值的形式
const beef, two, c = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
    Monday, Tuesday, Wednesday = 1, 2, 3
    Thursday, Friday, Saturday = 4, 5, 6
)
//枚举，iota从0开始,每当 iota 在新的一行被使用时，它的值都会自动加 1，每遇到一次 const 关键字，iota 就重置为 0
const (
    a = iota
    b
    c
)