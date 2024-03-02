## go编码规范

1. 命名规范

变量，常量，函数名使用驼峰法进行命名

2. error类型的全局变量请添加err或Err

```go
// 正常情况
var(
  ErrTest = errors.New("test")  // 可导出
  errTest = errors.New("test")  // 不可导出
)
```

3. 禁止接收者使用java,python等语言中的惯用命名

```go
// 不允许情况
type student struct{
  Name string
}
func(this *student)Name(){
  fmt.Println(this.Name)
}
```

4. 包名全小写，并且不要下划线

```go
// 不允许情况
Bao1
bao_1
```

5. 文件规范

- 文件名字使用小写字母
- 每行代码不能超过150个左右字符串（太长阅读很麻烦）
- go.mod和go.sum都要提交到代码库
- 同个struct方法放在同一个文件中
- pkg的功能描述建议写到doc.go文件中

6. 语言规范

- 禁止在if,for中对bool类型中进行等值判断
- 当函数以else结尾时，直接return

```go
func test(){
  if true {
    fmt.Println("test")
  }
  return 
}
```

- error类型返回一定要返回在末尾
- 有错误返回一定要处理（很重要）
- 参数包含context.Context时，总是作为第一个参数

7. 风格规范

- 使用tab进行缩进，并统一格式化
- 导出类型的一定要注释
- import需要按照标准库，第三方库，项目自身库的顺序分组排列







