package Utils

import "io/fs"

func GetDirNames(dirList []fs.FileInfo) []string {
	res := make([]string, 0)
	for _, dir := range dirList {
		if dir.IsDir() {
			res = append(res, dir.Name())
		}
	}
	return res
}

func GetFileNames(dirList []fs.FileInfo) []string {
	res := make([]string, 0)
	for _, file := range dirList {
		if !file.IsDir() {
			res = append(res, file.Name())
		}
	}
	return res
}
