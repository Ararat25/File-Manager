package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
	"sync"
	"net/http"
	"log"
	"bytes"
	"encoding/json"
)

const (
	sizeConversionFactor = 1000
)

func pathHandle(w http.ResponseWriter, r *http.Request) {
	root := r.URL.Query().Get("root")
	sort := r.URL.Query().Get("sort")

	if root == "" || sort == "" {
		http.Error(w, "Ошибка: пропущены нужные флаги.", http.StatusBadRequest)
		return
	}

	if !(sort == "ASK" || sort == "DESC") {
		http.Error(w, "Ошибка: флаг сорт не может быть с таким значением.", http.StatusBadRequest)
		return
	}

	output := outputFileProperty(root, sort)

	resp, err := json.Marshal(output)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusOK)

    w.Write(resp)

    w.Write(resp)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/path", pathHandle)

	err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Println(err)
		return
    }
}


func outputFileProperty(dirName string, sortMethod string) string {
	propertiesFiles, err := getPropertiesFiles(dirName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	pairs := sortFiles(propertiesFiles, sortMethod)

	var buffer bytes.Buffer

	w := tabwriter.NewWriter(&buffer, 0, 1, 3, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Name\tType\tSize")

	for _, pair := range pairs {
		fmt.Fprintf(w, "%s\t%s\t%s\n", pair.Key, pair.Value.fileType, formatSize(pair.Value.size))
	}

	w.Flush()

	return buffer.String()
}

var nameSizes = [5]string{"b", "Kb", "Mb", "Gb", "Tb"}

func getFlag() (string, string) {
	rootPtr := flag.String("root", "", "Путь до нужной директории")
	sortPtr := flag.String("sort", "ASK", "Параметр сортировки (возрастание/убывание)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование: go run fs.go --root=<путь_до_нужной_директории> --sort=<параметр_сортировки>\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *rootPtr == "" || *sortPtr == "" {
		fmt.Println("Ошибка: пропущены нужные флаги.")
		flag.Usage()
		return "", ""
	}

	if !(*sortPtr == "ASK" || *sortPtr == "DESC") {
		fmt.Println("Ошибка: флаг сорт не может быть с таким значением.")
		flag.Usage()
		return "", ""
	}

	dirName := *rootPtr
	sortMethod := *sortPtr

	return dirName, sortMethod
}

// formatSize преобразует размер в байтах к удобному для чтения виду
func formatSize(size int64) string {
	s := float64(size)

	if s < sizeConversionFactor {
		return fmt.Sprintf("%.2f B", s)
	}

	i := 0

	for s >= sizeConversionFactor && i < len(nameSizes)-1 {
		s /= float64(sizeConversionFactor)
		i++
	}

	return fmt.Sprintf("%.1f %s", s, nameSizes[i])
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

func processFile(fPath string, fileInfo fs.FileInfo, dirName string, file fs.DirEntry, propertiesFiles map[string]fileProperty, wg *sync.WaitGroup, mu *sync.Mutex) {
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
	propertiesFiles[file.Name()] = fileProperty{fileType: typeFile, size: size}
}

// getPropertiesFiles возвращает свойства файлов из заданной директории
func getPropertiesFiles(dirName string) (map[string]fileProperty, error) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	propertiesFiles := map[string]fileProperty{}

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
			go processFile(fPath, fileInfo, dirName, file, propertiesFiles, &wg, &mu)
		} else {
			processFile(fPath, fileInfo, dirName, file, propertiesFiles, &wg, &mu)
		}
		
	}

	wg.Wait()

	return propertiesFiles, nil
}

type fileProperty struct {
	fileType string
	size     int64
}

type Pair struct {
	Key   string
	Value fileProperty
}

// sortFiles cортирует мапу с размерами файлов по размеру
func sortFiles(propertiesFiles map[string]fileProperty, sortMethod string) []Pair {
	pairs := []Pair{}

	for key, value := range propertiesFiles {
		pairs = append(pairs, Pair{Key: key, Value: value})
	}

	if sortMethod == "ASK" {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Value.size < pairs[j].Value.size
		})
	}

	if sortMethod == "DESC" {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Value.size > pairs[j].Value.size
		})
	}

	return pairs
}
