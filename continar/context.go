package continar

type GoContext struct {
	msg string
}

var gctx = NewGoContext()

func NewGoContext() *GoContext {
	return &GoContext{
		msg: "",
	}
}

func SetMsg(msg string) {
	gctx.msg = msg
}

func GetMsg() string {
	return gctx.msg
}

func GetGoContext() *GoContext {
	return gctx
}
