package main
import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func gcd(a,b int) int{
    // fix negative numbers according to task
    if a<0 {
        a*=-1
    }
    if b<0 {
        b*=-1
    }

    // first special case: both numbers are same
    if a==b {
        return a
    }
    
    // second special case (most common when algorithm is done)
    if b==0 {
        return a
    }
    if a==0 {
        return b
    }
    
    // do the math
    r:=a%b
    
    // and finally run gcd recursively
    return gcd(b,r)
}

func lcm (a,b int) int64 {
    g:=gcd(a,b)
    return int64(a/g)*int64(b/g)*int64(g)
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Calculator fod gcd and lcm")
    fmt.Println("---------------------")
    for {
        var A,B int 
        var err error
        fmt.Printf("A= ")
        text, _ := reader.ReadString('\n')
        // convert CRLF to LF
        text = strings.Replace(text, "\n", "", -1)
        A,err=strconv.Atoi(text)
        if err!=nil{
            fmt.Printf("\nPlease enter only positive integer numbers or 0 to leave programm!\n")
            continue
        }
        if (A<1) {
           break
        }
        
        fmt.Printf("B= ")
        text, _ = reader.ReadString('\n')
        // convert CRLF to LF
        text = strings.Replace(text, "\n", "", -1)
        B,err=strconv.Atoi(text)
        if err!=nil{
            fmt.Printf("\nPlease enter only positive integer numbers or 0 to leave programm!\n")
            continue
        }
        if (B<1) {
            break
        }
    
        fmt.Printf("gcd=%d  lcm=%d\n\n",gcd(A,B),lcm(A,B))
    }
}
