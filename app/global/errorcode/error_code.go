package errorcode

// newErrorCode 錯誤代碼格式
type newErrorCode struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

var errorCode = map[string]newErrorCode{
	/** 共同 **/
	"SUCCESS":                 {1, "SUCCESS"},                    // 呼叫API成功
	"PERMISSION_DENIED":       {403, "PERMISSION DENIED"},        // 權限不足
	"CREATE_DIR_ERROR":        {1001, "CREATE DIR ERROR"},        // 建立資料夾失敗
	"GET_UPLOAD_FILE_ERROR":   {1002, "GET UPLOAD FILE ERROR"},   // 取得上傳檔案失敗
	"CREATE_FILE_ERROR":       {1003, "CREATE FILE ERROR"},       // 建立檔案失敗
	"GET_UPLOAD_TYPE_ERROR":   {1004, "GET UPLOAD TYPE ERROR"},   // 取得上傳類型錯誤
	"JSON_MARSHAL_ERROR":      {1005, "JSON MARSHAL ERROR"},      // json encode 失敗
	"JSON_UNMARSHAL_ERROR":    {1006, "JSON UNMARSHAL ERROR"},    // json decode 失敗
	"CHANGE_PARAMS_TYPE_FAIL": {1007, "CHANGE PARAMS TYPE FAIL"}, // 資料轉型失敗
	"PARSE_TIME_ERROR":        {1008, "PARSE TIME ERROR"},        // 時間格式轉換錯誤
	"VAILDATE_PARAMS_FAIL":    {1009, "VAILDATE PARAMS FAIL"},    // 參數型態驗證失敗
	"IMAGES_TOO_LARGE":        {1010, "IMAGES TOO LARGE"},        // 檔案過大
	"BIND_PARAMS_FAIL":        {1011, "BIND PARAMS FAIL"},        // 組合參數失敗
	"CRYPTION_ERROR":          {1012, "CRYPTION ERROR"},          // 密碼加密錯誤
	"GET_TIME_ZONE_ERROR":     {1013, "GET TIME ZONE ERROR"},     // 取當前時區錯誤
	"LOG_ID_NOT_EXIST":        {1014, "LOG ID NOT EXIST"},        // Log 身份證

	/** DB 錯誤 **/
	"DB_CONNECT_ERROR": {2000, "DB CONNECT ERROR"}, // DB連線失敗

	/** Redis 錯誤 **/
	"REDIS_CONNECT_ERROR":     {3000, "REDIS CONNECT ERROR"},     // Redis連線失敗
	"REDIS_INSERT_ERROR":      {3001, "REDIS INSERT ERROR"},      // Redis寫入失敗
	"REDIS_DELETE_ERROR":      {3002, "REDIS DELETE ERROR"},      // Redis刪除失敗
	"REDIS_APPEND_ERROR":      {3003, "REDIS APPEND ERROR"},      // Redis增加值失敗
	"REDIS_SET_EXPIRE_ERROR":  {3004, "REDIS SET EXPIRE ERROR"},  // Redis設定過期時間失敗
	"REDIS_CHECK_EXIST_ERROR": {3005, "REDIS CHECK EXIST ERROR"}, // 檢查Redis值是否存在時發生錯誤
	"REDIS_PING_ERROR":        {3006, "REDIS PING ERROR"},        // Redis Ping 錯誤
	"REDIS_GET_VALUE_ERROR":   {3007, "REDIS GET VALUE ERROR"},   // Redis 取值錯誤

	/** 呼叫 API 錯誤 **/
	"CURL_CREATE_FAIL":  {4000, "CURL CREATE FAIL"},  // CURL建立失敗
	"CURL_GET_FAIL":     {4001, "CURL GET FAIL"},     // CURL GET 失敗
	"CURL_POST_FAIL":    {4002, "CURL POST FAIL"},    // 取API失敗
	"API_CONNECT_ERROR": {4003, "API CONNECT ERROR"}, // 對外連線失敗
	"API_STATUS_ERROR":  {4004, "API STATUS ERROR"},  // 對外連線回傳code異常

	/** 其他 **/
	"AUTH_VAILDATE_FAIL": {5000, "AUTH VAILDATE FAIL"}, // 登入驗證失敗
}