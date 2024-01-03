func removeDuplicates(inputStream chan string, outputStream chan string){
    var st string
    for{
        b,ok := <- inputStream
        if !ok {
            break
        }
        if b != st{
            outputStream <- b
            st = b
        }  
    }
    defer close(outputStream)
}