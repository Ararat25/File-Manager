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
	sizeConversionFactor = 1000 // значение для конвертации размеров файлов
)

var (
	nameSizes = [5]string{"b", "Kb", "Mb", "Gb", "Tb"} // массив с значениями названий для размера файлов
)

// File структура для хранения свойств файлов
type File struct {
	Name     string `json:"Name"`     // название файла
	FileType string `json:"FileType"` // тип файла
	Size     string `json:"Size"`     // размер файла в форматированном виде
	ByteSize int64  `json:"-"`        // размер файла в байтах
}

const (
	ASC  = "asc"  // значение для сортировки по возрастанию
	DESC = "desc" // значение для сортировки по убыванию
)

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

// getFileProperty определяет свойства файла
func getFileProperty(fileInfo fs.FileInfo, file fs.DirEntry) File {
	typeFile := "file"
	size := fileInfo.Size()

	return File{Name: file.Name(), FileType: typeFile, Size: formatSize(size), ByteSize: size}
}

// getFileProperty определяет свойства директории
func getDirectoryProperty(dirPath string, dir fs.DirEntry) (File, error) {
	typeFile := "dir"
	var err error

	size, err := determineSize(dirPath)
	if err != nil {
		return File{}, err
	}

	return File{Name: dir.Name(), FileType: typeFile, Size: formatSize(size), ByteSize: size}, err
}

// setPropertiesFiles задает файлам из слайса соответствующие свойства
func setPropertiesFiles(dirName string) ([]File, error) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	var sliceFiles []File

	var wg sync.WaitGroup

	for i, file := range files {
		fPath := fmt.Sprintf("%s/%s", dirName, file.Name())

		fileInfo, err := os.Stat(fPath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		sliceFiles = append(sliceFiles, File{})

		if fileInfo.IsDir() {
			wg.Add(1)
			go func(curr int) {
				defer wg.Done()
				f, err := getDirectoryProperty(fPath, file)
				if err != nil {
					fmt.Println(err)
					return
				}
				(sliceFiles)[i] = f
			}(i)
		} else {
			f := getFileProperty(fileInfo, file)
			sliceFiles[i] = f
		}
	}

	wg.Wait()

	return sliceFiles, nil
}

// sortFiles сортирует слайс с файлами по размеру
func sortFiles(files []File, sortMethod string) []File {
	if sortMethod == ASC {
		sort.Slice(files, func(i, j int) bool {
			return files[i].ByteSize < files[j].ByteSize
		})
	}

	if sortMethod == DESC {
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
