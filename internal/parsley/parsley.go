package parsley

import "strings"

func MakeLines(s string) []string {
	return strings.Split(s, "\n")
}

func JoinLines(lines []string) string {
	return strings.Join(lines, "\n")
}

func ReplaceLastSubStr(s, old, new string) string {
	pos := strings.LastIndex(s, old)
	if pos == -1 {
		return s
	}
	return s[:pos] + new + s[pos+len(old):]
}

func GetFirstLine(s string) string {
	lines := MakeLines(s)
	if len(lines) == 0 {
		return s
	}
	return lines[0]
}

func GetLastLine(s string) string {
	lines := MakeLines(s)
	if len(lines) == 0 {
		return s
	}
	return lines[len(lines)-1]
}

func RemoveAllSubStr(s string, subs ...string) string {
	for _, sub := range subs {
		s = strings.ReplaceAll(s, sub, "")
	}
	return s
}

func CountLeadingSpaces(line string) int {
	count := 0
	for _, char := range line {
		if char != ' ' {
			break
		}
		count++
	}
	return count
}

func PrefixLines(str, prefix string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func FlattenLines(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t")
	}
	return lines
}

func FlattenStr(str string) string {
	lines := MakeLines(str)
	flat := FlattenLines(lines)
	return strings.Join(flat, "")
}

func TrimLeadingSpaces(str string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " ")
	}
	return strings.Join(lines, "\n")
}

func SliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func BackTick() string {
	return "`"
}

func ReplaceFirstLine(input, newLine string) string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 {
		lines[0] = newLine
	}
	return strings.Join(lines, "\n")
}

func ReplaceLastLine(input, newLine string) string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 {
		lines[len(lines)-1] = newLine
	}
	return strings.Join(lines, "\n")
}

func Squeeze(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
