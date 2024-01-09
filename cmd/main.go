package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Nahbox/L1-wb/internal/task1"
	"github.com/Nahbox/L1-wb/internal/task10"
	"github.com/Nahbox/L1-wb/internal/task11"
	"github.com/Nahbox/L1-wb/internal/task12"
	"github.com/Nahbox/L1-wb/internal/task13"
	"github.com/Nahbox/L1-wb/internal/task14"
	"github.com/Nahbox/L1-wb/internal/task15"
	"github.com/Nahbox/L1-wb/internal/task16"
	"github.com/Nahbox/L1-wb/internal/task17"
	"github.com/Nahbox/L1-wb/internal/task18"
	"github.com/Nahbox/L1-wb/internal/task19"
	"github.com/Nahbox/L1-wb/internal/task2"
	"github.com/Nahbox/L1-wb/internal/task20"
	"github.com/Nahbox/L1-wb/internal/task21"
	"github.com/Nahbox/L1-wb/internal/task22"
	"github.com/Nahbox/L1-wb/internal/task23"
	"github.com/Nahbox/L1-wb/internal/task24"
	"github.com/Nahbox/L1-wb/internal/task25"
	"github.com/Nahbox/L1-wb/internal/task26"
	"github.com/Nahbox/L1-wb/internal/task3"
	"github.com/Nahbox/L1-wb/internal/task4"
	"github.com/Nahbox/L1-wb/internal/task5"
	"github.com/Nahbox/L1-wb/internal/task6"
	"github.com/Nahbox/L1-wb/internal/task7"
	"github.com/Nahbox/L1-wb/internal/task8"
	"github.com/Nahbox/L1-wb/internal/task9"
)

func main() {
	var taskNum, solutionNum int

	for {
		fmt.Print("Введите номер задания (1-26): ")
		fmt.Scan(&taskNum)

		solutionsCount, err := countFilesInDirectory(fmt.Sprintf("internal/task%d", taskNum))
		if err != nil {
			log.Fatal(err)
		}

		if solutionsCount > 1 && taskNum != 21 {
			fmt.Printf("Доступных решений для задания №%d: %d\n", taskNum, solutionsCount)
			fmt.Printf("Введите номер решения (1-%d): ", solutionsCount)
			fmt.Scan(&solutionNum)
		}

		switch taskNum {
		case 1:
			task1.Task1()
		case 2:
			if solutionNum == 1 {
				task2.Task2_1()
			}
			if solutionNum == 2 {
				task2.Task2_2()
			}
		case 3:
			if solutionNum == 1 {
				task3.Task3_1()
			}
			if solutionNum == 2 {
				task3.Task3_2()
			}
		case 4:
			task4.Task4()
		case 5:
			if solutionNum == 1 {
				task5.Task5_1()
			}
			if solutionNum == 2 {
				task5.Task5_2()
			}
		case 6:
			task6.Task6()
		case 7:
			task7.Task7()
		case 8:
			task8.Task8()
		case 9:
			task9.Task9()
		case 10:
			task10.Task10()
		case 11:
			if solutionNum == 1 {
				task11.Task11_1()
			}
			if solutionNum == 2 {
				task11.Task11_2()
			}
		case 12:
			task12.Task12()
		case 13:
			task13.Task13()
		case 14:
			task14.Task14()
		case 15:
			task15.Task15()
		case 16:
			task16.Task16()
		case 17:
			task17.Task17()
		case 18:
			task18.Task18()
		case 19:
			task19.Task19()
		case 20:
			task20.Task20()
		case 21:
			task21.Task21()
		case 22:
			task22.Task22()
		case 23:
			task23.Task23()
		case 24:
			task24.Task24()
		case 25:
			task25.Task25()
		case 26:
			if solutionNum == 1 {
				task26.Task26_1()
			}
			if solutionNum == 2 {
				task26.Task26_2()
			}
			if solutionNum == 3 {
				task26.Task26_3()
			}
		}

		fmt.Println()

		taskNum = 0
		solutionNum = 0
		solutionsCount = 0
	}
}

func countFilesInDirectory(dirPath string) (int, error) {
	var fileCount int

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Если это файл (а не директория), увеличиваем счетчик
		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return fileCount, nil
}
