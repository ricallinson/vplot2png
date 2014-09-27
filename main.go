package main 

import (
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png" 
    "os"
) 

func plot(data *os.File) *image.Gray {
    // Work out the size from the source vplot file.
    img := image.NewGray(image.Rect(0, 0, 100, 100)) 
    for x := 20; x < 80; x++ { 
        y := x/3 + 15 
        img.Set(x, y, color.Black) 
    }
    return img
}

func main() { 

    flag.Parse()
    vplot := flag.Arg(0)
    dest := flag.Arg(1)

    if vplot == "" {
        fmt.Println("A source vplot file must be provide as the first argument.")
        return
    }

    if dest == "" {
        fmt.Println("A destination file must be provide as the second argument.")
        return
    }

    vplotData, err := os.Open(vplot)

    if err != nil {
        fmt.Println("Could not open the source vplot file.")
        return
    }

    destData := plot(vplotData)

    file, _ := os.Create(dest)
    defer file.Close()
    png.Encode(file, destData)
} 
