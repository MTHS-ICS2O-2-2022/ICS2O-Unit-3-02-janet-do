// Created by: Janet
// Created on: Sep 2020
//
// This program finds the volume of a pyramid

package main

import "fmt"

func main() {
    // This function finds the volume of a pyramid

    var length float64
    var width float64
    var height float64
    var volume float64

    // input
    fmt.Println("This program finds the volume of a pyramid.")
    fmt.Println()
    fmt.Print("Enter the length (in cm): ")
    fmt.Scanln(&length)
    fmt.Print("Enter the width (in cm): ")
    fmt.Scanln(&width)
    fmt.Print("Enter the height (in cm): ")
    fmt.Scanln(&height)
    fmt.Println()

    // process
    volume = (length * width * height) / 3.0

    // output
    fmt.Printf("The volume of the pyramid is: %.2f cm³\n", volume)
    fmt.Println()
    fmt.Println("Done.")
}
