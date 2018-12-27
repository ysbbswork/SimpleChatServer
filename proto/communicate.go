package proto

//Message 定义的总的通信结构
//整个包体包含个命令字头和数据包体
type Message struct {
	Cmd  int `json:"cmd"`
	Data string `json:"data"`
}


//命令字
const (
	Register = iota	//注册用户
	Login			//登录
	Msg			//消息
)

const (
	Success = 0
	Faild = 1
	UserAlreadyExist = 2
	UserNotExist =2
)

type RegisterReq struct{
	Name string
	Password string
}

type RegisterRsp struct{
	Flag int // 0成功 1失败 2重名了
}


type LoginReq struct{
	Uin uint
	Name string
	Password string
}

type LoginRsp struct {
	Flag int
	//flag为0登录成功，flag为密码错误，为2没有这个用户需要注册
}


type MessageReq struct{
	Msg string
}
type MessageRsp struct{
	Msg string
}

//用户status
const (
	UserOnline  = 1
	UserOffline = 2
)
