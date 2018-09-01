package utils

import "strings"

func NoHtml(str string) string {
	strScriptLeft := strings.Replace(str, "<script", "&lt;script", -1)
	strScriptRight := strings.Replace(strScriptLeft, "script>", "script&gt;", -1)
	return strings.Replace(strScriptRight, "\r\n", "\n", -1)
}
