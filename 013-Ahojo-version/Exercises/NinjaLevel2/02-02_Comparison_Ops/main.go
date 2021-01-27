package main 


import "fmt"


func main(){
    a := 42 == 42
    b:= 33 <= 44
    c:= "Hello" >= "World"
    d := b != a
    e := 44 < 44
    f := 5555 > 4444

    fmt.Println(a, b,c,d,e,f)
}

