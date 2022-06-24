package service

type Item struct {
	Ip       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
}
type Args struct {
	Key   string `json:"key"`
	Value Item   `json:"value"`
}

type ExecuteParams struct {
	Param map[string]Args `json:"args"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse() Result {
	return Result{
		Code: 200,
		Msg:  "Success.",
		Data: nil,
	}
}

const TOKENKEY = "token"

var Db *DB

func Init(db *DB) {
	Db = db
}

// 解析客户端的请求
func Parse(args ExecuteParams) (resp interface{}, err error) {

	for k, item := range args.Param {
		switch k {
		case "get":
			var val Item
			err := Db.GetWithCodec([]byte(item.Key), &val)
			return val, err
		case "set":
			err := Db.SetWithCodec([]byte(item.Key), item.Value)
			return nil, err
		case "del":
			err := Db.Del([]byte(item.Key))
			return nil, err
		case "clear":
			err := Db.Clear()
			return nil, err
		case "all":
			return Db.All()
		default:
		}
	}
	return
}
