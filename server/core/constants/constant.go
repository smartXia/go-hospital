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
	"南京",
	"上海",
	"北京",
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
