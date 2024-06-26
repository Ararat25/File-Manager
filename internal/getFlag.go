package internal

import (
	"flag"
	"fmt"
	"os"
)

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
