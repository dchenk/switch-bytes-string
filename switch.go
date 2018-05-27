package switch_bytes_string

const (
	strShort  = "abcd"
	strMedium = "abcdefgabcdefgabcdefgabcdefgabcde"
	strLong   = "abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga"
)

var Output string

func Str(s string) string {
	switch s {
	case "abc":
		return "---"
	case "abcdefgabcdefgabcdefgabcdefgabcde":
		return "---------"
	case "abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga":
		return "------"
	default:
		return "-----"
	}
}

func StrConst(s string) string {
	switch s {
	case strShort:
		return "---"
	case strMedium:
		return "---------"
	case strLong:
		return "------"
	default:
		return "-----"
	}
}

func Bts(s []byte) string {
	switch string(s) {
	case "abc":
		return "---"
	case "abcdefgabcdefgabcdefgabcdefgabcde":
		return "---------"
	case "abcdefgabcdefgabcdefgabcdefgabcdegabcdefgabcdefdbcdefgabcdefbcdefgabcdefegabcdefdegabcdefgagabcdefgabcdefgagabcdefgabcdefga":
		return "------"
	default:
		return "-----"
	}
}

func BtsConst(s []byte) string {
	switch string(s) {
	case strShort:
		return "---"
	case strMedium:
		return "---------"
	case strLong:
		return "------"
	default:
		return "-----"
	}
}
