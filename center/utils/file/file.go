package file

import "os"

// WriteFile 写文件
func WriteFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
