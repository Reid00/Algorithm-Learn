package stack

import (
	path2 "path"
	"strings"
)

// simplifyPath 简化路径
func simplifyPath(path string) string {

	stack := make([]string, 0)

	for _, v := range strings.Split(path, "/") {

		switch v {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")

}

// simplifyPath 简化路径
func simplifyPath2(path string) string {
	return path2.Clean(path)
}
