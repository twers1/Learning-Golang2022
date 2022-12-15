// Task 1 Задание 1 

res := 0
for i := 0; i < 10; i += 3 {
  res += i
} 
fmt.Println(res) // 18

// Task 2 Задание 2 (The input is the name of the programming language. You need to read it and output it.) На вход подается название языка программирования. Вам нужно считать его и вывести.
 
package main

import "fmt"

func main() {
    var input string
    fmt.Scanln(&input)
    
    fmt.Println(input)
}

// Task 3 Задание (Three integers are given as input. You need to count and print their sum and product on different lines.) На вход подаются три целых числа. Необходимо сосчитать и вывести их сумму и произведение на разных строках.

package main

import "fmt"

func main() {
    var a int
    var b int
    var c int
    
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    
    fmt.Println(a + b + c)
    fmt.Println(a * b * c)
}

// Functions Функции
/*
Functions are an important part of the Go language.

They allow you to define a block of code, name it, and call it in code. This makes the function code reusable, since you can call it several times in different parts of your program.

We've already seen some functions in our previous examples.
For example, the Println() function outputs text:

fmt.Println("Hello!")
Println is the name of the function, and "Hello!" - is the argument we pass to the function using parentheses.

Remember, to call a function, we always use the function name followed by parentheses.

Функции это важная часть языка Go.

Они позволяют вам определить блок кода, дать ему имя и вызвать его в коде. Это делает код функции многоразовым, так как вы можете вызывать его несколько раз в разных частях вашей программы.

Мы уже встречали некоторые функции в наших предыдущих примерах.
Например, функция Println() выводит текст:

fmt.Println("Привет!")
Println - это имя функции, а "Привет!" - это аргумент, который мы передаем функции с помощью круглых скобок.

Помните, чтобы вызвать функцию, мы всегда используем имя функции, за которым следуют круглые скобки.
*/

// Task 4 Задание 4 You need to write a line of code that prints the text "Test!" using the Println() function. Вам нужно написать строку кода, которая с помощью функции Println() выведет текст "Тест!"

fmt.Println("Тест!")

/*
In addition to the functions available in Go packages, you can create your own functions!
Use the keyword func to create your own function. 

В дополнение к функциям, доступным в пакетах Go, вы можете создавать свои собственные функции!
Используйте ключевое слово func, чтобы создать свою собственную функцию. 
*/

func welcome() {
  fmt.Println("Привет!")
}

// We can now call our function in the program inside the main main function main: Use our function name, followed by parentheses - welcome()
// Теперь мы можем вызвать нашу функцию в программе внутри главной функции main:  Используйте имя нашей функции, за которым следуют круглые скобки - welcome()

package main

import "fmt"

func welcome() {
  fmt.Println("Hello")
}

func main() {
  welcome()
  welcome()
} 


// Task 5 Задание 5 (example)
package main

import "fmt"

func line() {
    fmt.Println("-----")
}

func main() {
    line()
}

// Arguments Аргументы

// Arguments are a way to take input parameters for our functions.
// Аргументы - это способ принимать входные параметры для наших функций.

func welcome(name string) {
  fmt.Println("Привет, "+name)
} 

package main

import "fmt"

func welcome(name string) {
  fmt.Println("Привет, " + name)
}

func main() {
  welcome("Максим")
  welcome("Андрей")
}

// Task 6 Задание 6 

func convert(x int) {
  fmt.Println(x/100)
} 
func main() {
  convert(300)
} // 3

// For a function to accept multiple arguments, separate them with commas.
// Чтобы функция принимала несколько аргументов, разделяйте их запятыми.

func sum(a int, b int) {
  fmt.Println(a + b)
}

// Task 7 Задание 7

func mult(x, y int) {
    fmt.Println(x*y)
}

// Operator return Оператор return

// The operator return terminates the function and returns the value provided to the calling code.
// Оператор return завершает функцию и возвращает предоставленное значение вызвавшему ее коду.

func sum(x, y int) int {
  return x + y
}

// Task 8 Задание 8

func concat(a, b string) string {
    return a + b
}

/* 
A useful feature of Go is that functions can return multiple values!
The swap() function takes two integer arguments and returns them in swapped order.
Note that the type of each return value must be declared: in our case it is (int, int).

Полезной особенностью Go является то, что функции могут возвращать несколько значений!
Функция swap() принимает два целочисленных аргумента и возвращает их в поменявшемся порядке.
Обратите внимание, что тип каждого возвращаемого значения должен быть объявлен: в нашем случае это (int, int).
*/
func swap(x, y int) (int, int) {
    return y, x
} 

// or или 

a, b := swap(42, 8)

// Defer 

/* 
The defer operator ensures that the function will only be called after the surrounding function returns.
The above code will first print "Hey!" and only then will it print the result of welcome().
This is because the welcome() call is delayed, that is, it waits until main() has finished executing, and only then is it called.

Оператор defer гарантирует, что функция будет вызвана только после того, как вернется окружающая функция.
Приведенный выше код сначала выведет "Эй!" и только после этого выведет результат функции welcome().
Это происходит потому, что вызов функции welcome() отложен, то есть она ждет, пока main() закончит выполнение, и только после этого происходит её вызов.
*/

package main

import "fmt"

func welcome() {
  fmt.Println("Добро пожаловать")
}

func main() {
  defer welcome()
  fmt.Println("Эй!")
}

// If you put off several function calls, they will be executed in order. The defer calls are stacked on top of each other, so they are executed in last-in-first-out (LIFO) order.
// Если вы отложили несколько вызовов функций, они будут выполняться в порядке очереди. Вызовы defer`ов укладываются друг на друга по принципу стека (или стопки), поэтому они выполняются в порядке "последний-пришел-первый-вышел" (LIFO).

package main

import "fmt"

func main() {
  fmt.Println("начало")

  for i := 0; i < 5; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("конец")
}



