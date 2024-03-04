// Package les4: Измените программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды
package les4

import (
	"fmt"
	"os"
)

func les4() {
	var s, sep string
	for i := 1; i <= len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "      // сепаратор или разделитель между аргументами если их больше 2
		fmt.Println(s) //ln вывод на новую строку
	}
}
