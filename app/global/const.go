package global

// DurianMa db master
const DurianMa = "durian_master"

// DurianSl db slave
const DurianSl = "durian_slave"

/** 正則規則 **/
const (
	AdminAccount  = `[a-zA-Z0-9_]{2,30}` // AdminAccount 管理者帳號
	AdminPassword = `[a-zA-Z0-9]{6,20}`  // AdminPassword 管理者密碼
)

/** Redis Cache Time **/
const (
	RedisCyExpire = 300 // CY API資料 5分鐘
	RedisDBExpire = 600 // 資料庫資料 10分鐘
)

/** 平台設定 **/
const (
	FrontEnd = "frontend" // 前台
	BackEnd  = "backend"  // 後台
)

/** 類別清單 **/
const (
	Marquee = 1 // 跑馬燈
	News    = 2 // 最新報導
)

/** 錯誤 Log 類型 **/
const (
	WarnLog  = "Warn"  // 警告型 Log
	FatalLog = "Fatal" // 嚴重型 Log
)
