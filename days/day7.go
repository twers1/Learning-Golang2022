// Task 1 Задание 1 The product of all Произведение всех
package main

import "fmt"

func main() {
    var x int
    fmt.Scanln(&x)

    y := 1   
    for i := 1; i <= x; i++ {
        y *= i
    }

    fmt.Println(y)
}

// Signposts Указатели

// Pointers are special variables that store addresses of values in memory. In Go we declare a pointer with * Here p is a pointer to an integer value..: 
// Указатели - это специальные переменные, которые хранят адреса значений в памяти. В Go мы объявляем указатель с помощью символа * Здесь p - это указатель на целочисленное значение.:

var p *int

/* 
We know how to declare a pointer, but how do we assign it an address in memory? 
To do this, we use the & operator, which returns the address of the variable in memory.

Мы знаем, как объявить указатель, но как присвоить ему адрес в памяти? 
Для этого используется оператор &, который возвращает адрес переменной в памяти.

Now p is a pointer and stores the address x in memory.
Let's print its value:

Теперь p является указателем и хранит в памяти адрес x.
Давайте выведем его значение:

If we want to access the base value of a pointer, we can use the * operator
Если мы хотим получить доступ к базовому значению указателя, мы можем использовать оператор *
*/ 

x := 42
p := &x

fmt.Println(p) // выведет что-то вроде 0xc000100010

fmt.Println(*p) // выведет 42

// The * operator can also be used to change the value at the memory address pointed to by the pointer:
// Оператор * также может быть использован для изменения значения по адресу в памяти, на который указывает указатель:

x := 42
p := &x

*p = 8
fmt.Println(*p) // выведет 8
fmt.Println(x)  // выведет 8

// The * operator can also be used to change the value at the memory address pointed to by the pointer:
// Мы смогли изменить значение x с помощью указателя p. Оператор * называется оператором разыменования.
