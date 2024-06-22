package system

// code
type WxResponse struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}
type TokenRes struct {
	Token string `json:"access_token"`
}

// 手机code
type PhoneResponse struct {
	ErrCode   int       `json:"errcode"`
	ErrMsg    string    `json:"errmsg"`
	PhoneInfo PhoneInfo `json:"phone_info"`
}
type Watermark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       Watermark `json:"watermark"`
}

// token
type TokenResponse struct {
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	CurrentTime int64     `json:"current_time"`
	EventID     string    `json:"event_id"`
	Data        TokenData `json:"data"`
}

type TokenData struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	Role  int    `json:"role"`
	Token string `json:"token"`
}

// 用户列表
type WeixinUserData struct {
	Total      int      `json:"total"`
	Count      int      `json:"count"`
	Data       UserData `json:"data"`
	NextOpenID string   `json:"next_openid"`
}

type UserData struct {
	OpenID []string `json:"openid"`
}

// 用户详细信息
type WeixinUserInfo struct {
	Subscribe      int    `json:"subscribe"`
	OpenID         string `json:"openid"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	HeadImgURL     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionID        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	TagIDList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QRScene        int    `json:"qr_scene"`
	QRSceneStr     string `json:"qr_scene_str"`
}
