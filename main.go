package main 

import (
    "image" 
    "image/color"
    "image/png" 
    "os"
) 

func main() { 
    img := image.NewRGBA(image.Rect(0, 0, 100, 100)) 
    for x := 20; x < 80; x++ { 
        y := x/3 + 15 
        img.Set(x, y, color.Black) 
    }
    file, _ := os.Create("test.png")
    defer file.Close()
    png.Encode(file, img)
} 
