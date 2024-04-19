//                          uy ishi-1
// package main

// import (
// 	"fmt"
// )

// func factCalc(num int, ch chan int) {
// 	fact := 1
// 	for i := 2; i <= num; i++ {
// 		fact *= i
// 	}
// 	ch <- fact
// }

// func main() {
// 	numbers := []int{5, 6, 7, 8, 9}
// 	results := make(chan int)
// 	for _, num := range numbers {
// 		go factCalc(num, results)
// 	}
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Printf("Faktoriali %d: %d\n", numbers[i], <-results)
// 	}
// 	close(results)
// }





//                                 uy ishi-2
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	files := []string{"file1.txt", "file2.txt", "file3.txt"}
// 	var mergedContent []byte
// 	for _, filename := range files {
// 		file, err := os.Open(filename)
// 		if err != nil {
// 			fmt.Printf("Xatolik: %v\n", err)
// 			continue
// 		}
// 		defer file.Close()

// 		fileInfo, err := file.Stat()
// 		if err != nil {
// 			fmt.Printf("Xatolik: %v\n", err)
// 			continue
// 		}

// 		fileSize := fileInfo.Size()
// 		content := make([]byte, fileSize)

// 		_, err = file.Read(content)
// 		if err != nil {
// 			fmt.Printf("Xatolik: %v\n", err)
// 			continue
// 		}

// 		mergedContent = append(mergedContent, content...)
// 	}
// 	fmt.Println(string(mergedContent))
// 	err := os.WriteFile("merged_file.txt", mergedContent, os.ModePerm)
// 	if err != nil {
// 		fmt.Printf("Xatolik: %v\n", err)
// 		return
// 	}
// 	fmt.Println("Faylga muvaffaqiyatli yozildi: merged_file.txt")
// }


