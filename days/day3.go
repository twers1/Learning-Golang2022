// Comparison operators Операторы сравнения

// Comparison operators are used to compare two values and return a bool value as the result: true if the comparison condition is met, and false if it is not.
// Операторы сравнения используются для сравнения двух значений и возвращают в качестве результата значение типа bool: true, если условие сравнения выполняется, и false, если не выполняется.

x := 42
y := 8
    
fmt.Println(x == y) 
    
fmt.Println(x != y) 

fmt.Println(x > y)  

fmt.Println(x < y)  

/* 
The AND (or &&) logical operator returns true only when both conditions are true.
The OR (or ||) logical operator returns true when one (or both) of the conditions are true.
The NOT (or !) logical operator returns true when the condition is false (essentially reversing the result of the condition). 

Логический оператор AND (или &&) возвращает истину только тогда, когда оба условия истинны.
Логический оператор OR (или ||) возвращает истину, когда одно (или оба) из условий истинно.
Логический оператор NOT (или !) возвращает истину, если условие ложно (по сути, меняя результат условия на противоположный). 
*/

x := 42
y := 8
    
fmt.Println(x!=y && x<=y)  

fmt.Println(x!=y || x<=y)  

fmt.Println(!(x>y))        

// User input Ввод пользователя

/* 
The "fmt" package also allows you to accept input from the user of the program.
To accept input, you must use the Scanln() function and provide it with a variable, which must contain the input value: 

The input variable will now contain the value that the user enters when the program starts.

Пакет "fmt" также позволяет принимать ввод от пользователя программы.
Чтобы принять ввод, нужно использовать функцию Scanln() и предоставить ей переменную, которая должна содержать вводимое значение: 

Теперь переменная input будет содержать значение, которое пользователь введет при запуске программы.
*/

var input string
fmt.Scanln(&input)

fmt.Println(input)

// In this case the & symbol returns the variable address
// В данном случае символ & возвращает адрес переменной

// Task 1 (The program reads two integers from the input and outputs their sum) Задание 1 (программа считывала два целых числа со входа и выводила их сумму)

package main

import "fmt"

func main() {
    var a int
    var b int
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Println(a + b)
}


// Conditional operator IF Условный оператор IF

// The if statement can be used to make decisions and execute code when a given condition is true.
// Оператор if можно использовать для принятия решений и выполнения кода, когда заданное условие истинно.

x := 42

if x > 18 {
  fmt.Println("Разрешено")
}

// Task 2 (so it outputs "Pass" if the score variable is greater than or equal to 75.) Задание 2 (чтобы он выводил "Проходи", если переменная score будет больше или равна 75.)

package main

import "fmt"

func main() {
    var score = 80
    if score >= 75 {
        fmt.Println("Проходи")
    }
}

// The else operator can be used to execute code if the condition of the if statement is false.
// Оператор else можно использовать для выполнения кода, в том случае, если условие оператора if ложно.

x := 14

if x > 18 {
  fmt.Println("Разрешено")
} else {
  fmt.Println("Не разрешено")
}

/* 
Sometimes a variable is only needed for if/else statements.
To do this, the if statement in Go can start with a variable declared before the condition: 
Note the semicolon after the variable declaration - it is used to separate expressions.

Variables declared in an if statement are only available in if/else blocks.

Иногда переменная нужна только для операторов if/else.
Для этого оператор if в Go может начинаться с объявления переменной перед условием: 
Обратите внимание на точку с запятой после объявления переменной - она используется для разделения выражений.

Переменные, объявленные в операторе if, доступны только в блоках if/else.
*/

if x := 42; x > 18 {
   fmt.Println("Разрешено")
} else {
   fmt.Println("Не разрешено")
}
