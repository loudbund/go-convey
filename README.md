# go-convey
数据判断简易封装，主要功能
1. 判断未知变量的预计值，
2. 判断map格式[或json]指定深度的值
3. 判断一组未知变量的值

可用于校验函数的返回值，或api接口返回值，从而判断函数或转接口是否正确等场所

## 安装
go get github.com/loudbund/go-convey

## 引入
```golang
import "github.com/loudbund/go-convey/convey_v1"
```

## 使用
1. 使用前需要先获取实例 Convey := convey_v1.NewConvey()
2. 执行相等判断，不等判断，json相等判断等
3. 获取判断结果 numSucc,numFail:=Convey.GetNum()
4. 重置判断结果 Convey.ResetNum()
5. 下一轮判断
```golang
// 代码参见example.go
```