package exception

//业务错误码常量
const (
	SUCCESE  = 0
	ERROR    = -1
	UNDEFINE = 1
	//501-1000 数据库错误
	ERROR_DATA_QUERY_INVALID = 501
	//1001-2000 文章模块错误
	ERROR_ARTICAL_NOT_EXIST = 2001
	//2001-3000 分类模块错误
	ERROR_CATEGORY_USED      = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002
	//3001-4000 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_NOT_EXIST  = 1005
	ERROR_TOKEN_OVERDUE    = 1006
	ERROR_TOKEN_WRONG      = 1007
	ERROR_TOKEN_TYPE_WRONG = 1008
)

var codemap = map[int]string{
	SUCCESE: "OK",
	ERROR:   "FAIL",
	//501-1000 数据库错误
	ERROR_DATA_QUERY_INVALID: "数据库查询不合法",
	//1001-2000 文章模块错误
	ERROR_ARTICAL_NOT_EXIST: "文章不存在",
	//2001-3000 分类模块错误
	ERROR_CATEGORY_USED:      "分类已被使用",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
}

type Err struct {
	Code int
	Msg  string
}

func (this *Err) Error() string {
	if this.Code == UNDEFINE {
		return this.Msg
	}
	return codemap[this.Code]
}
