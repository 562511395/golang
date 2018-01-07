//变量声明，当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil；变量也遵循可见性原则,有作用域,函数体外声明的变量为全局变量
//一般写法
var a, b *int


var a int
var b bool
var str string

//并行写法
var (
    a int
    b bool
    str string
)
//自动推断变量类型，全局
var (
    HOME = os.Getenv("HOME")
    USER = os.Getenv("USER")
    GOROOT = os.Getenv("GOROOT")
)
//自动推断变量类型，局部
a := 1
//vartest.go 通过runtime包在运行时获取所在的操作系统类型，以及如何通过 os 包中的函数 os.Getenv() 来获取环境变量中的值，并保存到 string 类型的局部变量 path 中