package convey_v1

import (
	"errors"
	"github.com/loudbund/go-json/json_v1"
	. "github.com/smartystreets/goconvey/convey"
)

// 1、结构体 -------------------------------------------------------------------------
type XConvey struct {
	numSucc int      // 成功次数
	numFail int      // 失败次数
	errMsg  []string //
}

// 2、全局变量 -------------------------------------------------------------------------

// 3、初始化函数 -------------------------------------------------------------------------

// 4、开放函数 -------------------------------------------------------------------------

// 实例化
func NewConvey() *XConvey {
	return &XConvey{
		numSucc: 0,
		numFail: 0,
		errMsg:  make([]string, 0),
	}
}

// 重置次数
func (Me *XConvey) ResetNum() {
	Me.numSucc = 0
	Me.numFail = 0
}

// 获取成功失败次数
func (Me *XConvey) GetNum() (numSucc, numFail int) {
	return Me.numSucc, Me.numFail
}

// 获取错误信息
func (Me *XConvey) GetErrs() []string {
	return Me.errMsg
}

// 断言1：相等才正确
func (Me *XConvey) ShouldEqual(actual interface{}, expected ...interface{}) error {
	if result := ShouldEqual(actual, expected...); result == "" {
		Me.numSucc++
		return nil
	} else {
		Me.numFail++
		Me.errMsg = append(Me.errMsg, result)
		return errors.New(result)
	}
}

// 断言2：不等才正确
func (Me *XConvey) ShouldNotEqual(actual interface{}, expected ...interface{}) error {
	if result := ShouldNotEqual(actual, expected...); result == "" {
		Me.numSucc++
		return nil
	} else {
		Me.numFail++
		Me.errMsg = append(Me.errMsg, result)
		return errors.New(result)
	}
}

// 断言3：json相等才正确
func (Me *XConvey) JShouldEqual(actual interface{}, keys []interface{}, expected interface{}) (interface{}, error) {
	Data, err := json_v1.GetJsonInterface(actual, keys...)
	if err != nil {
		Me.numFail++
		Me.errMsg = append(Me.errMsg, err.Error())
		return nil, err
	}
	// data判断
	if err := Me.ShouldEqual(Data, expected); err != nil {
		Me.errMsg = append(Me.errMsg, err.Error())
		return nil, err
	}
	return Data, nil
}

// 断言
func (Me *XConvey) so(Name string, actual interface{}, assert Assertion, expected ...interface{}) {
	// Actual := ""
	// Len := 100
	// if _, ok := actual.(string); ok {
	// 	Actual = actual.(string)
	// } else {
	// 	A, _ := json_v1.JsonEncode(actual)
	// 	Actual = A
	// }
	// if len(Actual) < 100 {
	// 	Len = len(Actual)
	// }
	//
	if result := assert(actual, expected...); result == "" {
		Me.numSucc++
		// Me.msg("so", Name, "success", "", "", "", Actual[:Len], 0)
	} else {
		Me.numFail++
		// Me.GroupFailNum++
		// Actual, _ := json_v1.JsonEncode(actual)
		// if len(Actual) < 100 {
		// 	Len = len(Actual)
		// }
		// Me.msg("so", Name, "fail", "", "", "result", Actual[:Len], 0)
	}
}

// json数据断言+返回
// func (Me *XConvey) soJson(Name string, actual interface{}, keys []interface{}, assert Assertion, expected ...interface{}) interface{} {
// 	Data, err := json_v1.GetJsonInterface(actual, keys...)
// 	Me.so("取数据err须为nil", err, ShouldEqual, nil)
// 	Me.so(Name, Data, assert, expected)
// 	return Data
// }
