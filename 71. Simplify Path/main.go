package main

import "regexp"

func simplifyPath(path string) string {

	if path[len(path)-1] != '/' {
		path += "/"
	}

	result := regexp.MustCompile("\\/{2,}").ReplaceAll([]byte(path), []byte("/"))

	for dl := 0; dl != len(result); {
		dl = len(result)
		result = regexp.MustCompile("\\/\\.\\/").ReplaceAll(result, []byte("/"))
	}

	for dl := 0; dl != len(result); {
		dl = len(result)
		result = regexp.MustCompile("\\/[^\\/]+\\/\\.{2}\\/").ReplaceAll(result, []byte("/"))
	}

	result = regexp.MustCompile("\\/\\.+\\/").ReplaceAll(result, []byte("/"))

	if result[len(result)-1] == '/' && len(result) > 1 {
		result = result[:len(result)-1]
	}

	return string(result)
}

func main() {
	//println(simplifyPath("/home//foo/../test/../another/"))
	//println(simplifyPath("/a/./b/../../c/"))
	println(simplifyPath("/a//b////c/d//././/.."))
}
