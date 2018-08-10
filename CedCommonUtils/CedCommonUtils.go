package CedCommonUtils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

// ReadLine 从文件中逐行读取
func ReadLine(filename string, handler func(line string)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		// line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// ReadBytes 从文件中读取数据
func ReadBytes(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetCurrentCodePath 获得现在所执行的文件所在的路径
func GetCurrentCodePath() (filePath string, dirPath string) {
	_, filePath, _, _ = runtime.Caller(1)
	dirPath = path.Dir(filePath)
	return
}
