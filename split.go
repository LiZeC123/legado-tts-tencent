package main

import (
	"regexp"
)

func SplitText(text string) []string {
	// 前置检查：整体长度不超过150直接返回
	if len([]rune(text)) <= 150 {
		return []string{text}
	}

	// 按中文标点进行语义分割
	pattern := regexp.MustCompile(`[^。！？\n]+[。！？\n]*`)
	paragraphs := pattern.FindAllString(text, -1)

	var (
		result []string
		buffer []rune // 改用rune切片提升性能
	)

	for _, p := range paragraphs {
		current := []rune(p)
		// 合并后长度允许则追加到缓冲区
		if len(buffer)+len(current) <= 150 {
			buffer = append(buffer, current...)
			continue
		}

		// 缓冲区有内容时先落盘
		if len(buffer) > 0 {
			result = append(result, string(buffer))
			buffer = buffer[:0] // 重用内存
		}

		// 处理超长段落
		if len(current) > 150 {
			result = append(result, splitByLength(current)...)
		} else {
			buffer = append(buffer, current...)
		}
	}

	// 处理最终缓冲区
	if len(buffer) > 0 {
		result = append(result, string(buffer))
	}

	return result
}

// 处理超长段落的分割
func splitByLength(r []rune) []string {
	var res []string
	for i := 0; i < len(r); i += 150 {
		end := i + 150
		if end > len(r) {
			end = len(r)
		}
		res = append(res, string(r[i:end]))
	}
	return res
}
