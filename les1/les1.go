package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ { // os.Args - это переменная типа слайс полуоткрытый
		//(первый элемент входит, последний исключается как в математике [)) и хранит в виде элементов полный путь к .exe
		//запускаемому файлу а дальше элементами являются входные аргументы к запускаемому файлу (например -port=8080 и -user=admin)
		//их можно ввести в комнадную строку PowerShell или через софт в вкладке edit configurations.. --> programm arguments
		//для создания статических не изменяемых аргументов к запускаемому проекту выгодно использовать конфигурационный файл json
		//для создания "гибких" аргументов к проекту выгодно использовать системные настройки?
		s += sep + os.Args[i]
		sep = " " // сепаратор или разделитель между аргументами если их больше 2
	}
	fmt.Println(s) //ln вывод на новую строку
}
