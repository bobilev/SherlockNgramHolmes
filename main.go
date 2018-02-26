package main

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"github.com/bobilev/SherlockNgramHolmes/util"
	"github.com/bobilev/SherlockNgramHolmes/config"
	"strconv"
	"math"
	"strings"
	"flag"
	"fmt"
	"time"
)
//record[0] - _time
//record[1] - src_user
//record[2] - src_ip
//record[3] - src_port
//record[4] - dest_user
//record[5] - dest_ip
//record[6] - dest_port
//record[7] - input_byte
//record[8] - output_byte
var FileNameString = flag.String("f","","path to file")
var MaxCountSort = flag.Int("s", 5, "sort int")
var Empty = flag.Bool("e",false,"empty")
var Vision = flag.Bool("v",false,"show information in the console")
func main() {
	start := time.Now()
	mapList1 := make(map[string]int)
	mapList2 := make(map[string]int)
	var arrayNgramST1 []string
	mapListRepetitive := make(map[string]map[string]int)
	mapListRepetitive2 := make(map[string]map[string]int)
	var arreyResultText [][]string
	flag.Parse()
	config.MaxCountSort = *MaxCountSort
	config.FileNameString = *FileNameString
	config.Empty = *Empty

	if config.MaxCountSort > 25 || config.MaxCountSort < 1 {
		config.MaxCountSort = 5
		fmt.Println("Sort больше 25 или меньше 1 нельзя (default: 5)")
	}
	if config.FileNameString == "" {
		fmt.Println("Error: not file")
		os.Exit(1)
	}

	// Load a csv file.
	f, errFile := os.Open(config.FileNameString)
	if errFile != nil {
		fmt.Println("Error: not file")
		os.Exit(1)
	}
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		src_user := record[1]
		src_ip := record[2]
		if config.Empty && src_user == ""{
			src_user = "nil"
		}
		if config.Empty && src_ip == ""{
			src_ip = "nil"
		}
		findTheBiggest(mapList1,src_user,1)
		inputByte , _:= strconv.Atoi(record[7])//Конвертируем строку в число
		findTheBiggest(mapList2,src_user,inputByte)

		line1 := src_ip+","+record[3]+","+record[4]+","+record[5]+","+record[6]+","+record[7]+","+record[8]
		line2 := src_user+","+record[3]+","+record[4]+","+record[5]+","+record[6]+","+record[7]+","+record[8]
		findRepetitive(mapListRepetitive,src_user,line1)
		findRepetitive(mapListRepetitive2,src_ip,line2)

		arrayNgramST1 = append(arrayNgramST1,record[1]+","+record[3]+","+record[5]+","+record[6])
	}
	fmt.Println("the file is now complete","time:",time.Since(start))
	mapListRepetitiveUnity1 := util.MergeTwoMap(mapListRepetitive)//Обединяем все в один map
	mapListRepetitiveUnity2 := util.MergeTwoMap(mapListRepetitive2)//Обединяем все в один map

	arreyResultText = append(arreyResultText,util.UtilSort(mapList1, config.MaxCountSort))
	arreyResultText = append(arreyResultText,util.UtilSort(mapList2, config.MaxCountSort))
	arreyResultText = append(arreyResultText,util.UtilSort(mapListRepetitiveUnity1, config.MaxCountSort))
	arreyResultText = append(arreyResultText,util.UtilSort(mapListRepetitiveUnity2, config.MaxCountSort))
	arreyResultText = append(arreyResultText,util.UtilSort(ngram(arrayNgramST1,3), config.MaxCountSort))
	arreyResultText = append(arreyResultText,util.UtilSort(ngram(arrayNgramST1,4), config.MaxCountSort))
	arreyResultText = append(arreyResultText,util.UtilSort(ngram(arrayNgramST1,5), config.MaxCountSort))

	var lineAnswer []string
	for key,val := range arreyResultText {
		lineAnswer = append(lineAnswer,config.AnswerMass[key])
		for key1, val1 := range val {
			nnn := ""
			if key > 3 {
				nnn = "\n"
			}
			lineAnswer = append(lineAnswer,"["+strconv.Itoa(key1+1)+"] "+nnn+val1)
		}
		lineAnswer = append(lineAnswer,"")
	}
	util.CreateFile(lineAnswer,"result")
	if *Vision {
		for _, val := range lineAnswer {
			fmt.Println(val)
		}
	}
	fmt.Println("time:",time.Since(start))
}
//Решение задачи 1-2
func findTheBiggest(mapList map[string]int, user string, index int) {
	if user != "" {
		if val, ok := mapList[user]; ok {//Если уже есть в списке
			mapList[user] = val + index
		} else {//Если еще нету в списке, то добавляем
			mapList[user] = index
		}
	}
}
//Решение задачи 3-4
func findRepetitive(mapList map[string]map[string]int,user string,line string) {
	if user != "" {
		if val, ok := mapList[user]; ok {//Если уже есть в списке
			if val1, ok1 := val[line]; ok1 {//Если запрос уже есть то плюсуем
				val[line] = val1 + 1
			} else {
				val[line] = 1
			}
		} else {//Если еще нету в списке, то добавляем
			mapList[user] = make(map[string]int)
			mapList[user][line] = 1
		}
	}
}
//Решение задачи 5
func ngram(words []string, size int)  map[string]int{
	countMapNrgam := make(map[string]int)
	offset := int(math.Floor(float64(size / 2)))

	for i, _:= range words {
		if i < offset || i+size-offset > len(words) {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], "\n")
		countMapNrgam[gram]++
	}
	return countMapNrgam
}