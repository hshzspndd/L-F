package services

type ResponseErrorForm struct {
	Code    int
	Message string
}

func (e *ResponseErrorForm) Errdor() string {
	return e.Message
}

var (
	ErrHashPassword  = &ResponseErrorForm{500, "密码加密失败"}
	ErrNoPermission  = &ResponseErrorForm{403, "没有权限"}
	ErrUserCheckFail = &ResponseErrorForm{500, "用户信息校验失败"}
	ErrWrongPassword = &ResponseErrorForm{403, "密码错误"}
	ErrUserNotFound  = &ResponseErrorForm{404, "用户不存在"}
	ErrUserExists    = &ResponseErrorForm{409, "用户已存在"}
	ErrDatabase      = &ResponseErrorForm{500, "数据库出错"}
)
