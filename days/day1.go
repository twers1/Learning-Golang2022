// Основы Basics

// Давайте начнем практиковаться с создания простой программы Go для вывода на экран слов "Go это весело!". 
// Let's start practicing by creating a simple Go program to display the words "Go is fun! 

package main // каждая программа на языке Go начинает выполняться в главном пакете, поэтому объявляет этот пакет, как главный(т.е. main)
// every Go program starts in the main package, so it declares this package as the main one

import "fmt" // мы можем импортировать и использовать в коде для выполнения различных задач, один из самых популярных это fmt - format(обеспечивает функциональность 
// ввода и вывода ;  we can import and use in code to perform various tasks, one of the most popular is fmt - format(provides functionality for input and output 

func main() { // имена функций импортируются из пакета Println() - это имя функции, которую мы импортируем из пакета, поэтому оно начинается с заглавной буквы.
// function names are imported from the package Println() is the name of the function we import from the package, so it starts with a capital letter.
    fmt.Println("Go это весело!")
}


/* Дополните (замените нижние подчеркивания) код, чтобы программа вывела на экран сообщение "Go отличный язык!".
 Complete (replace the lower underscores) the code so that the program displays the message "Go great language!
_______ main

import "___"

func main() {
    fmt._______("Go отличный язык!")
}

 Response: */

package main

import "fmt"

func main() {
    fmt.Println("Go отличный язык!")
}

/* Дополните код таким образом, чтобы получилась рабочая программа на Go, которая будет выводить текст "Второй модуль позади!"
Complete the code so that you have a working program in Go that outputs the text "The second module is behind you!"

package ____

______ "fmt"

func main() {
    ___.Println("Второй модуль позади!")
} 
Response: */

package main

import "fmt"

func main() {
    fmt.Println("Второй модуль позади!")
}

