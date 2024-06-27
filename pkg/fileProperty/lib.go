package fileProperty

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const (
	sizeConversionFactor = 1000
)

var nameSizes = [5]string{"b", "Kb", "Mb", "Gb", "Tb"}

type File struct {
	Name     string
	FileType string
	Size     string
	ByteSize int64 `json:"-"`
}

// OutputFileProperty возвращает слайс файлов со свойствами
func OutputFileProperty(dirName string, sortMethod string) ([]File, error) {
	propertiesFiles, err := setPropertiesFiles(dirName)
	if err != nil {
		return nil, err
	}

	pairs := sortFiles(propertiesFiles, sortMethod)

	return pairs, nil
}

// determineSize определяет полный размер директории вместе с файлами
func determineSize(f string) (int64, error) {
	var size int64

	err := filepath.Walk(f, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += info.Size()

		return nil
	})
	if err != nil {
		return 0, err
	}

	return size, nil
}

// processFile функция для записи файлов в слайс
func processFile(fPath string, fileInfo fs.FileInfo, file fs.DirEntry, files *[]File, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	typeFile := "file"
	size := fileInfo.Size()

	if file.IsDir() {
		typeFile = "dir"

		var err error

		size, err = determineSize(fPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	mu.Lock()
	defer mu.Unlock()
	*files = append(*files, File{Name: file.Name(), FileType: typeFile, Size: formatSize(size), ByteSize: size})
}

// setPropertiesFiles задает файлам из слайса соответствующие свойства
func setPropertiesFiles(dirName string) ([]File, error) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	var sliceFiles []File

	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	for _, file := range files {
		wg.Add(1)

		fPath := fmt.Sprintf("%s/%s", dirName, file.Name())

		fileInfo, err := os.Stat(fPath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if fileInfo.IsDir() {
			go processFile(fPath, fileInfo, file, &sliceFiles, &wg, &mu)
		} else {
			processFile(fPath, fileInfo, file, &sliceFiles, &wg, &mu)
		}

	}

	wg.Wait()

	return sliceFiles, nil
}

// sortFiles сортирует слайс с файлами по размеру
func sortFiles(files []File, sortMethod string) []File {
	if sortMethod == "ASK" {
		sort.Slice(files, func(i, j int) bool {
			return files[i].ByteSize < files[j].ByteSize
		})
	}

	if sortMethod == "DESC" {
		sort.Slice(files, func(i, j int) bool {
			return files[i].ByteSize > files[j].ByteSize
		})
	}

	return files
}

// formatSize преобразует размер файла в байтах к удобному для чтения виду
func formatSize(size int64) string {
	s := float64(size)

	if s < sizeConversionFactor {
		return fmt.Sprintf("%.1f B", s)
	}

	i := 0
	for s >= sizeConversionFactor && i < len(nameSizes)-1 {
		s /= float64(sizeConversionFactor)
		i++
	}

	return fmt.Sprintf("%.1f %s", s, nameSizes[i])
}
