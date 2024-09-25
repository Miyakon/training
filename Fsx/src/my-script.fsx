printfn "Hello World"

let add1 x = x + 1

let squarePlusOne  x = 
    let square = x * x
    square + 1

add1(1)
squarePlusOne(3)

let areEqual x y =
    (x = y)

