package main

import (
	"regexp"
	"strings"
)

var gTerminatorRegex *regexp.Regexp
var gPunctuationRegex *regexp.Regexp

func init() {
	// 定义中文终止符号的正则表达式
	// 包括：。？！；…（省略号）以及换行符
	terminators := `([。？！；\n]|……)+`

	// 定义中文标点符号的正则表达式
	punctuation := `^[\s。？！，、；："'‘’“”（）《》【】…·~\-—=+*%￥#@&]+$`

	// 编译正则表达式
	gTerminatorRegex = regexp.MustCompile(terminators)
	gPunctuationRegex = regexp.MustCompile(punctuation)
}

// SplitText 按照中文终止符号分割文本并清理结果
func SplitText(text string) []string {
	// 分割文本
	segments := gTerminatorRegex.Split(text, -1)

	// 清理结果
	var result []string
	for _, seg := range segments {
		trimmed := strings.TrimSpace(seg)
		if isMeaningfulSegment(trimmed) {
			result = append(result, trimmed)
		}
	}

	return result
}

// isMeaningfulSegment 检查片段是否有效（非空且不仅包含标点符号）
func isMeaningfulSegment(segment string) bool {
	if segment == "" {
		return false
	}

	return !gPunctuationRegex.MatchString(segment)
}
