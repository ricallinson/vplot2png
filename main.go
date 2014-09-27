package main 

import (
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png" 
    "os"
) 

func main() { 

    flag.Parse()
    vplot := flag.Arg(0)
    dest := flag.Arg(1)

    if vplot == "" {
        fmt.Println("A source vplot file must be provide as the first parameter.")
        return
    }

    if dest == "" {
        fmt.Println("A destination file must be provide as the second parameter.")
        return
    }

    // Work out the size from the source vplot file.
    img := image.NewRGBA(image.Rect(0, 0, 100, 100)) 
    for x := 20; x < 80; x++ { 
        y := x/3 + 15 
        img.Set(x, y, color.Black) 
    }

    file, _ := os.Create(dest)
    defer file.Close()
    png.Encode(file, img)
} 
