package types

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrorEmptyParam         = ErrJSON(1001, "缺少必要参数")
	ErrorParamInvalid       = ErrJSON(1002, "参数错误，请确认是否已传入参数且该参数类型符合要求")
	ErrorParams             = ErrJSON(1003, "不支持的参数值，请检查")
	ErrorIllegalCross       = ErrJSON(1093, "非法定义，标签解析越界")
	ErrorInvalidExtention   = ErrJSON(1095, "无效表达式")
	ErrorInvalidCompareExt  = ErrJSON(1096, "无效的比较运算")
	ErrorNotFoundValid      = ErrJSON(1097, "变量不存在")
	ErrorNotFoundData       = ErrJSON(1102, "没有数据")
	Error404PageNotFound    = ErrJSON(1109, "404 找不到该页面")
	ErrorTokenWrong         = ErrJSON(2006, "appKey或token验证失败")
	ErrorUserNotFound       = ErrJSON(2007, "用户不存在或密码错误")
	ErrorTokenInvalid       = ErrJSON(2008, "token无效")
	ErrorPermissionDenied   = ErrJSON(2009, "没有操作权限")
	ErrorInvalidTagString   = ErrJSON(2014, "无效标签定义")
	ErrorForTagClose        = ErrJSON(2015, "闭合标签有误，请检查")
	ErrorTimestamp          = ErrJSON(2020, "ts非法")
	ErrorCateNotFound       = ErrJSON(2021, "分类不存在")
	ErrorUploadFile         = ErrJSON(2023, "上传文件错误")
	ErrorCannotDelSelf      = ErrJSON(2024, "不能删除自己")
	ErrorPhoneNumberType    = ErrJSON(2025, "手机号格式错误")
	ErrorNickNameType       = ErrJSON(2026, "用户昵称长度只能在1~5个字之间")
	ErrorPasswordType       = ErrJSON(2026, "密码长度需在6~10位之间")
	ErrorSavingImage        = ErrJSON(2027, "保存图片失败")
	ErrorFunctionInvalid    = ErrJSON(2029, "函数未生效，因为你并没有提供足够的参数")
	ErrorNotFoundTemplate   = ErrJSON(2030, "找不到模板文件")
	ErrorNotArrayOfValue    = ErrJSON(2031, "in_array() 首位参数必需是数组形式")
	ErrorIfExtentions       = ErrJSON(2032, "if 表达式错误，请检查语法是否合法。多余的空格，或关键字书写错误，都将触发这个错误")
	ErrorExtention          = ErrJSON(2033, "不支持的表达式符号")
	ErrorWebTitleRepeat     = ErrJSON(2034, "存在名称相同的站点")
	ErrorWebURLRepeat       = ErrJSON(2035, "存在链接相同的站点")
	ErrorMysqlConfig        = ErrJSON(2036, "未检测到数据，请确保数据库已正确连接")
	ErrorMysqlConnection    = ErrJSON(2037, "无法连接到数据库，请检查配置信息是否正确")
	ErrorFrequentOperation  = ErrJSON(2038, "操作频繁，请先休息一会吧!")
	ErrorRepeatHandleLike   = ErrJSON(2039, "您已经点过了~!")
	ErrorLikeNowAllow       = ErrJSON(2040, "操作错误，您已对该信息点踩~!")
	ErrorUnLikeNowAllow     = ErrJSON(2041, "操作错误，您已对该信息点赞~!")
	ErrorArticleTitleExists = ErrJSON(2043, "失败，已存在相同的文章标题")
	ErrorCateTitleExists    = ErrJSON(2042, "失败，已存在相同的分类标题")
)

type JSON struct {
	Code int16       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ErrJSON(code int16, msg string, data ...interface{}) error {
	j := &JSON{Code: code, Msg: msg}
	if len(data) > 0 {
		j.Data = data[0]
	}
	buf, _ := json.Marshal(j)
	return errors.New(string(buf))
}

// WithError 在 {code:-1, msg:""} 的原有基本上为msg字段新增error msg
func WithError(err, msg error) error {
	var (
		ret JSON
		str = err.Error()
	)
	json.Unmarshal([]byte(str), &ret)
	if ret.Code > 0 && msg != nil {
		ret.Msg += "; reason: " + msg.Error()
		b, _ := json.Marshal(ret)
		return errors.New(string(b))
	}
	return nil
}

func GetErrorJSON(err error, data interface{}) []byte {
	if err == nil {
		b, _ := json.Marshal(&JSON{
			Msg:  "success",
			Data: data,
		})
		return b
	}

	var (
		ret JSON
		str = err.Error()
	)
	if strings.Contains(str, `"code":`) {
		if strings.Contains(str, `"data":`) {
			json.Unmarshal([]byte(str), &ret)
			ret.Data = data
			b, _ := json.Marshal(ret)
			return b
		}
		return []byte(str)
	}

	b, _ := json.Marshal(&JSON{Code: -1, Msg: err.Error()})
	return b
}
