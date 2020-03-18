package auth

type AccessLimit struct {
	AccessKey string
	ClientIp string
}

type LoginUser struct {
	Id   int64
	Type  string
	HasLogin bool
}
