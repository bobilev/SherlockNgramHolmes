package util

import (
	"sort"
	"strconv"
	"os"
	"fmt"
	"bufio"
)
//Здесь происходит сортировка по возрастанию
func UtilSort(m map[string]int, value int) []string{
	var ints []int
	var mm []string

	for _, val := range m {
		ints = append(ints,val)//Данные поступают в массив для будущей сортировки
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))//Непосредственная сортировка

	//Для сохранения правильной последовательности засовываем все в массив
	for i := 0; i < value ; i++ {
		key := findekey(m,ints[i])//Находим
		mm = append(mm,key+" - "+strconv.Itoa(ints[i]))
	}
	return mm
}
//Ищем ключь в изначальном Map по передаваемому значению
func findekey(m map[string]int, value int) string{
	key := ""
	for k, v := range m {
		if v == value {
			key = k
			delete(m, k);
			return key
		}
	}
	return key
}
//Переделываем двумерную Map в одномерную, для того чтобы пропустить ее через func UtilSort()
func MergeTwoMap(Map map[string]map[string]int) map[string]int{
	mapListRepetitiveUnity := make(map[string]int)
	for key, val := range Map {//Обединяем все в один map
		for key1, val2 := range val {
			mapListRepetitiveUnity[key+","+key1] = val2
		}
	}
	return mapListRepetitiveUnity
}

func CreateFile(lines []string, path string) error{
	f, err := os.Create(path+".txt")
	defer f.Close()
	if err != nil {
		fmt.Println("---------------[ERR] dont create file nameFile")
		fmt.Println(err)
	}
	w := bufio.NewWriter(f)
	for _, line := range lines {
		//Накапливаем все в буфер
		fmt.Fprintln(w, line)
	}
	fmt.Println("[File save]")

	return w.Flush()

}