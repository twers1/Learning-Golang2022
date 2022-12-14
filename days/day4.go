// Switch operator Оператор switch
num := 3

switch num {
  case 1:
    fmt.Println("Один")
  case 2:
    fmt.Println("Два")
  case 3:
    fmt.Println("Три")
  default:
    fmt.Println(num)
}

x := 2

switch {
  case x > 0 && x < 10:
    fmt.Println("something")
  case x > 10:
    fmt.Println("something else")
}

// Task 1 Задание 1 

package main

import "fmt"

func main() {
    var gender int
    switch gender {
        case 1:
        fmt.Println("Мужчина")
        case 2:
        fmt.Println("Женщина")
        default:
        fmt.Println("Еще не определился")
    }
}

/*
In other programming languages, such as C++ or Java, each case must have a break statement to stop the switch statement.
In Go, a break statement is not needed because switch cases are evaluated from top to bottom, stopping when the case condition is true.

В других языках программирования, таких как C++ или Java, каждый случай (case) должен иметь оператор break, чтобы остановить выполнение оператора switch.
В Go оператор break не нужен, так как случаи switch оцениваются сверху вниз, останавливаясь, когда условие случая верно.
*/

// Cycle for Цикл for 

/*
Loops are used to repeat code until a certain condition occurs.
The only loop construct in Go is for, which consists of three components: initialize counter, condition, change counter.

Циклы используются для повторения кода до тех пор, пока не наступит определенное условие.
Единственной конструкцией цикла в Go является for, который состоит из трех компонентов: инициализации счетчика, условия, изменения счетчика
*/

for i := 0; i < 5; i++ {
   fmt.Println(i)
}

// Task 2 Задание 2 
sum := 0
for i := 1; i <= 3; i++ {
  sum += i
}
fmt.Println(sum) // 6

// Task 3 Задание 3 

package main

import "fmt"

func main() {
  var num int 
  fmt.Scanln(&num)
  res :=num*2
  fmt.Println(res)
}

// Task 4 Задание 4

package main

import "fmt"

func main() {
    for i := 9; i > 0; i-- {
    fmt.Println(i)
    }
}

// Task 5 Задание 5 

x := 7
res := 0
switch {
case x > 8:
  res += 1
case x > 3 && x < 10:
  res += 2
case x > 5:
  res += 3
}
fmt.Println(res) // 2

// Task 6 Задание 6

package main

import "fmt"

func main() {
    height := 195
    if height > 185 {
        fmt.Println("Да")
    } else {
        fmt.Println("Нет")
    }
}
