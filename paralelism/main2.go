func task2(c chan string, st string){
    st = st + " "
    c <- st
    c <- st
    c <- st
    c <- st
    c <- st
}