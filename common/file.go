package common

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

// 监听文件新增内容
func WatchFileNewLine(filename string, ch chan string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取当前文件的大小
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// 移动文件指针到文件末尾
	_, err = file.Seek(fileSize, 0)
	if err != nil {
		return err
	}

	// 开始监听文件变化
	for {
		// 获取当前文件的大小
		fileInfo, err := file.Stat()
		if err != nil {
			klog.Error(err)
			continue
		}
		currentSize := fileInfo.Size()

		// 如果文件大小有变化，说明有新数据写入
		if currentSize > fileSize {
			// 读取新增的数据
			newData := make([]byte, currentSize-fileSize)
			_, err := file.Read(newData)
			if err != nil && err != io.EOF {
				klog.Error(err)
				continue
			}

			// 更新文件大小
			fileSize = currentSize

			data := string(newData)
			for _, line := range strings.Split(data, "\n") {
				ch <- line
			}
		}
		// 每隔一段时间检查文件变化
		time.Sleep(time.Second)
	}
}

// 创建文件
func CreateFile(filePath string, content string) error {
	if _, err := os.Create(filePath); err != nil {
		return err
	}
	return os.WriteFile(filePath, []byte(content), 0644)
}

// 获取文件大小
func GetFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	size := fileInfo.Size()
	return size, nil
}

// 文件大小可视化
func FileSizeToView(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	}
	if size < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(size)/1024)
	}
	if size < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(size)/(1024*1024))
	}
	if size < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(size)/(1024*1024*1024))
	}
	return fmt.Sprintf("%.2fTB", float64(size)/(1024*1024*1024*1024))
}

// 获取某个目录下所有文件,不包含目录
func GetAllFiles(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path.Base(filepath))
		}
		return nil
	})
	return files, err
}
