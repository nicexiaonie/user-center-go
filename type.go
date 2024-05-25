package user_center_go

type Request struct {
	RequestId string `json:"request_id"` // 请求ID
	TraceId   string `json:"trace_id"`   // 链路ID
	Source    string `json:"source"`     // Source
	Secret    int    `json:"secret"`     // 1: 加密. 接口会进行解密获取请求数据  2: 未加密, 原始获取, 返回信息脱敏
	Body      any    `json:"body"`
}

// step1 获取用户基础信息
type ResponseUserBaseInfo struct {
	Code    int         `json:"code"`
	Message any         `json:"msg"`
	Body    UserInfoRes `json:"body"`
}

type UserInfoRes struct {
	UserId       uint64 `json:"user_id"`
	AccountName  string `json:"account_name"`
	NickName     string `json:"nick_name"`
	ProfilePhoto string `json:"profile_photo"`
	Password     string `json:"password"` // 密码
	Email        string `json:"email"`
	IsRealAuth   int    `json:"is_real_auth"`
	RealName     string `json:"real_name"`
	IdCard       string `json:"id_card"`
	RegisterTime string `json:"register_time"`
	Gender       int    `json:"gender"`
	BirthdayDay  string `json:"birthday_day"`
	PhoneNumber  string `json:"phone_number"`
	OpenId       string `json:"open_id"`
}

// step 发送登录短信
type ResponseServiceSmsSendLogin struct {
	Code    int                  `json:"code"`
	Message any                  `json:"msg"`
	Body    _ServiceSmsSendLogin `json:"body"`
}
type _ServiceSmsSendLogin struct {
	ID          uint64 ` json:"id"`
	SmsType     int    ` json:"smsType"`     // 短信类型,  1：登录验证
	Project     string ` json:"project"`     // 项目
	PhoneNumber string ` json:"phoneNumber"` // 手机号码
	UserID      uint64 ` json:"userId"`      // 用户ID
	Code        string ` json:"code"`        // 验证码
	ValidSecond uint32 ` json:"validSecond"` // 有效时间(秒)
}

// 登录
type ApiUserLoginReq struct {
	LoginType    string `json:"login_type"`     // 登录方式： sms:短信登录, wechat:微信登录 one_click:一键登录 activity:活动 password:密码
	UserId       int64  `json:"user_id"`        // 用户识别码
	Password     string `json:"password"`       // 密码
	PhoneNumber  string `json:"phone_number"`   //手机号码
	SmsCode      string `json:"sms_code"`       // 短信验证码
	WechatOpenId string `json:"wechat_open_id"` // 微信openId, 此字段兼容存在不建议使用
	WechatCode   string `json:"wechat_code"`    // 微信code
	BirthTime    string `json:"birth_time"`
	// step 以下为附加信息
	Ip                 string `json:"ip"`
	AppVersion         string `json:"app_version"`
	MachineCode        string `json:"machine_code"`
	PhoneType          string `json:"phone_type"`
	PhoneVersion       string `json:"phone_version"`
	PhoneSystemVersion string `json:"phone_system_version"`
}

// step修改信息
type ApiSetBaseInfo struct {
	UserId       uint64 `json:"user_id"`
	AccountName  string `json:"account_name"`
	NickName     string `json:"nick_name"`
	ProfilePhoto string `json:"profile_photo"`
	Gender       int    `json:"gender"`
	BirthdayDay  string `json:"birthday_day"`
}

// step 绑定微信 同步微信信息
type ApiBindWeChatReq struct {
	UserId          uint64 `json:"user_id"`            // 唯一不可变 user主表主键ID
	WechatCode      string `json:"wechat_code"`        // 微信授权code
	UserWXCacheData any    `json:"user_wx_cache_data"` // 兼容需废弃
}

// step 实名认证
type ApiRealNameReq struct {
	UserId   uint64 `json:"user_id"`   // 唯一不可变 user主表主键ID
	RealName string `json:"real_name"` // 真实姓名
	IdCard   string `json:"id_card"`   // 身份证号码
}

// step 获取用户组织关系
type ApiGetTreeUserReq struct {
	UserId uint64 `json:"user_id"` // 唯一不可变 user主表主键ID
}

// step 获取用户组织关系返回
type ResponseTreeUserInfo struct {
	Code    int              `json:"code"`
	Message any              `json:"msg"`
	Body    ResponseTreeUser `json:"body"`
}

// step 获取用户组织关系返回二级
type ResponseTreeUser struct {
	OrgID    uint32 `json:"org_id"`
	ParentID uint32 `json:"parent_id"`
	OrgName  string `json:"org_name"`
	Perm     string `json:"perm"`
}

// step 绑定用户组织关系
type ApiBindTreeUserReq struct {
	UserId uint64 `json:"user_id" binding:"required"` // 唯一不可变 user主表主键ID
	OrgID  uint32 `json:"org_id"  binding:"required"`
}

// step 绑定用户组织关系返回
type ResponseBindTreeUseInfo struct {
	Code    int         `json:"code"`
	Message any         `json:"msg"`
	Body    interface{} `json:"body"`
}

// step 解绑定用户组织关系
type ApiUnBindTreeUserReq struct {
	UserId uint64 `json:"user_id" binding:"required"` // 唯一不可变 user主表主键ID
	OrgID  uint32 `json:"org_id"  binding:"required"`
}

// step 解绑定用户组织关系返回
type ResponseUnBindTreeUserInfo struct {
	Code    int         `json:"code"`
	Message any         `json:"msg"`
	Body    interface{} `json:"body"`
}

// step 获取组织下级
type ApiGetyTreeUserChildrenReq struct {
	ID uint32 `json:"id" binding:"required"`
}

// step 获取组织下级返回
type ResponseGetyTreeUserChildrenInfo struct {
	Code    int              `json:"code"`
	Message any              `json:"msg"`
	Body    TreeUserChildren `json:"body"`
}

type TreeUserChildren struct {
	Children []TreeUser `json:"children"`
}

type TreeUser struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	ParentID uint32 `json:"parent_id"`
}
