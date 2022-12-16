// The scope Область видимости (scope)

/*
The scope is where a variable can be used.
In Go, there are two main scopes: local and global.
A variable defined in a function is called a local variable. Their scope is only in the body of the function. This means that they only exist within their function.

Область видимости - это то, где переменная может быть использована.
В языке Go существует две основных области видимости: локальная (local) и глобальная (global).

Переменная, определенная в функции, называется локальной переменной. Их область видимости находится только в теле функции. Это означает, что они существуют только в пределах своей функции.
*/

func test() {
  var x = 8
  fmt.Println(x)
}

package main

import "fmt"

func test() {
  var x = 8
  fmt.Println(x)
}

func main() {
  fmt.Println(x)
}
 
// выведет ошибку undefined: x

// Task 1 Задание 1 

var x = 13
func change() {
  x += 1
}
func set(x int) {
  x += 3
}
func main() {
  change()
  set(x)
  fmt.Println(x)
}

// Task 2 Задание 2 

func calc(a int) (int, int) {
  return a + 2, a + 1 
} 

func main() {
  x, y := calc(5)
  fmt.Println(x + y)
} // 13

// Task 3 Задание 3 (Write a max() function that should take as input two integers and return the larger of the two integers.) Напишите функцию max() которая должна принимать на вход два целых числа и возвращать большее из них.

func max(x, y int) int {
    if x>y {
        return x
    } else {return y}
}
