// Task 1 Задание 1 

package main 
 
import "fmt" 
 
func main() { 
    var money int 
 
    fmt.Scan(&money) 
 
    if money > 1000 { 
        fmt.Println("Apple") 
    } else if money >= 500 && money <= 1000 { 
        fmt.Println("Samsung") 
    } else { 
        fmt.Println("Nokia с фонариком") 
    } 
}

// Task 2 Задание 2

a := 4
p := &a
a += 2
*p = *p - 1
fmt.Println(a) // 5

// Passing pointers to functions Передача указателей в функции

// We can pass pointers as parameters to a function. Мы можем передавать указатели в качестве параметров функции.

package main

import "fmt"

func change(val int) {
  val = 8
}
func change_ptr(ptr *int) {
  *ptr = 8
}

func main() {
  x := 42

  change(x)
  fmt.Println(x) // выведет 42

  change_ptr(&x)
  fmt.Println(x) // выведет 8
}

/*
The change() function takes an integer argument and changes its value.
The change_ptr() function does the same, using a pointer.

By running the code you will see that change() has not changed the value of our variable x, because the argument is just a copy of its value, while change_ptr() has changed the actual value of x, because it uses an address in memory as an argument.

Функция change() принимает целочисленный аргумент и изменяет его значение.
Функция change_ptr() делает то же самое, используя указатель.

Выполнив код, вы увидите, что функция change() не изменила значение нашей переменной x, потому что аргумент является просто копией ее значения, в то время как функция change_ptr() изменила фактическое значение x, потому что в качестве аргумента использовала адрес в памяти.
*/
