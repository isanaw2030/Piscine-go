package main

import "os"

func main() {
    
   
    for i := 'a'; i <= 'z'; i++ {
       
        os.Stdout.Write([]byte{byte(i)})
    }

  
    os.Stdout.Write([]byte{'\n'})
}


