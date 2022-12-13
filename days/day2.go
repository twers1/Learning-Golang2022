// Variable Переменные

// Go uses the var keyword to declare variables. В языке Go для объявления переменных используется ключевое слово var. 
var i int

/*
The above code declares a variable named i of type int.
Int is short for integer and is a type of variable that can hold integers. We can assign a value to the variable and print it out: 

В приведенном выше коде объявлена переменная с именем i типа int.
int это сокращение от integer (целое число) и представляет собой тип переменной, в которой могут храниться целые числа. Мы можем присвоить переменной значение и вывести его:*/
var i int = 8
fmt.Println(i) 

/* Task 1 Задание 1 
Complete the program code so that it creates a variable x and assigns it the value 42, then outputs this variable.
Допишите код программы, чтобы она создала переменную x и присвоила ей значение 42, после чего вывела эту переменную.
*/
package main

import "fmt"

func main() {
    var x int = 42
    fmt.Println(x)
}

// Если вы присваиваете переменной значение, объявление типа можно опустить, так как Go автоматически примет тип присвоенного значения:
// If you assign a value to a variable, you can omit the type declaration because Go will automatically take the type of value assigned: 

var i, j = 8, 42

// Go supports short variable declaration using the operator :=
// Go поддерживает короткое объявление переменных с использованием оператора :=
var i int = 8
fmt.Println(i) 

// Data types Типы данных

/*
float32 - single-precision floating-point value.
float64 - double-precision floating point value.
string - text value.
bool - boolean value true or false.

float32 - значение с плавающей точкой одинарной точности.
float64 - значение с плавающей точкой двойной точности.
string - текстовое значение.
bool - булево значение true или false.
 */

var x int = 42
var y float32 = 1.37
var name string = "James"
var online bool = true // if no boolean value is specified, it defaults to false  если не указано булевое значение, то по умолчанию оно является false 

fmt.Println(name)  
fmt.Println(x)      
fmt.Println(y)      
fmt.Println(online) 

// Constants Константы

// Constants are declared as variables, but with the keyword const, and must be immediately assigned a value: (here, pi is a constant that cannot be changed)
// P.S. Constants cannot be declared using the syntax :=

// Константы объявляются как переменные, но с ключевым словом const, и им необходимо сразу присвоить значение: (здесь, pi - это константа, которую нельзя изменить)
// P.S. Константы не могут быть объявлены с использованием синтаксиса :=

const pi = 3.14

// Arithmetic operators Арифметические операторы

// Go supports all common arithmetic operators.
// Go поддерживает все распространенные арифметические операторы.

x := 42
y := 8
      
res := x + y
fmt.Println(res) 
      
res = x - y
fmt.Println(res)  

res = x * y
fmt.Println(res)  

res = x / y
fmt.Println(res)  
      
res = x % y
fmt.Println(res)  

// The + (plus) operator can also be used to join strings.
// Оператор + (плюс) также может быть использован для объединения строк.

x := "привет "
y := "мир"
fmt.Println(x + y) // выведет "привет мир"
