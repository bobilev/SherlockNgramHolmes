# SherlockNgramHolmes #

Установка Golang
---------

Читайте [https://golang.org/doc/install](https://golang.org/doc/install), я верю в вас =)


Команды
-------
	* -f - путь к файлу
	* -s - Число указывающее в каком количестве получить результат (по умолчанию: 5)
	* -e - Учитывать пустые src_user/src_ip при анализе (по умолчанию не учитывает)
	* -v - Выводить информацию в консоль (по умолчанию не выводит)

Пример
------
	go run main.go -f file

или

	go build
	./SherlockNgramHolmes -f file
	./SherlockNgramHolmes -f file -s 10 -e -v




# P.S #
* Чтение файла занимает 13 сек на ssd kingston - 480/480 Мб/с (так и не смог прикрутить многопоточность тут)
* Анализ занимает примерно 10 сек

* В задании 3/4 в формулировке "Поиск регулярных запросов (запросов выполняющихся периодически)" акцент на  регулярно/периодически можно понимать по-разному. Во первых это можно понять как часто встречающиеся тела запроса, следовательно он регулярен и периодически попадается. Во вторых это можно понять как появление этих запросов с определенной частотой по времени. То есть тут нужно писать алгоритм выявления временных закономерностей. Я в выбрал реализацию первого варианта.

* В 5 задании при поиске N-gram я заметил что в файле последовательность записи лога по времени перемешана. В теории это может повлиять на конечный результат, и коллокации могут получится другими.