package wechat

type CodeMsg struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type AccessTokenResp struct {
	CodeMsg
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type SendTextReq struct {
	ToUser                 string `json:"touser"`
	ToParty                string `json:"toparty"`
	ToTag                  string `json:"totag"`
	MsgType                string `json:"msgtype"`
	AgentID                string `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableIDTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
	Text                   struct {
		Content string `json:"content"`
	} `json:"text"`
}

type SendTextResp struct {
	CodeMsg
	InvalidUser    string `json:"invaliduser"`
	InvalidParty   string `json:"invalidparty"`
	InvalidTag     string `json:"invalidtag"`
	UnlicensedUser string `json:"unlicenseduser"`
	MsgID          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}
