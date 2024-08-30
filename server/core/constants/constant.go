package constants

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)
const (
	AUTH   = AUTHStr
	ADD    = ADDStr
	EDIT   = EDITStr
	QUERY  = QUERYStr
	IMPORT = IMPORTStr
	REPORT = REPORTStr
	UPLOAD = UPLOADStr
)
const (
	AUTHStr   = "auth"
	ADDStr    = "add"
	EDITStr   = "edit"
	QUERYStr  = "query"
	IMPORTStr = "import"
	REPORTStr = "export"
	UPLOADStr = "upload"
)
const (
	WinOs = "windows"
	Mac   = "mac"
	LunOS = "lunix"
)

var Diqupaihang = []string{
	"河北省", "山西省", "辽宁省", "吉林省", "黑龙江省", "江苏省", "浙江省", "安徽省", "福建省",
	"江西省", "山东省", "河南省", "湖北省", "湖南省", "广东省", "海南省", "四川省", "贵州省",
	"云南省", "陕西省", "甘肃省", "青海省", "天津市", "上海市", "重庆市", "内蒙古自治区",
	"广西壮族自治区", "西藏自治区", "宁夏回族自治区", "新疆维吾尔自治区", "北京市", "天津市", "上海市", "重庆市",
	"内蒙古自治区", "广西壮族自治区", "西藏自治区", "宁夏回族自治区", "新疆维吾尔自治区", "香港特别行政区", "澳门特别行政区",
}
var Duilienianling = []string{
	"幼儿",
	"儿童",
	"少年",
	"青年",
	"中年",
	"老年",
}
var Duilieyanzhongxing = []string{
	"轻",
	"轻度温和",
	"温和",
	"温和严重",
	"建议手术",
	"必须手术",
}
var Duiliefenlei = []string{
	"脊柱侧弯",
	"矢状面障碍",
	"椎体滑脱",
	"长短脚",
	"背疼",
	"其它",
}
var ActionEvent = []string{
	"register",
	"point",
}

const (
	REGISTER = "register"
	POINT    = "point"
)

var ActionEventChange = map[string]int{
	REGISTER: 5,
	POINT:    10,
}
var ActionEventChangeCn = map[string]string{
	REGISTER: "完成信息提交",
	POINT:    "完成打卡提交",
}
