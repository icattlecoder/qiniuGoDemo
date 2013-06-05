package main
import(
	"bufio"
	"strings"
	"os"
	"fmt"
)
func main(){
	f, _ := os.Open("./qrsb.keys")
	br := bufio.NewReader(f)
	line := 1
	for{
		l, _, err:= br.ReadLine()
		if err!=nil{
			break
		}
		i:= len(strings.Split(string(l),"\t"))
		if i!=4{
			fmt.Println("err!line=",line)
			os.Exit(1)
		}else{
			fmt.Println("line=",line,"OK!")
		}
		line++
	}
	fmt.Println("OK")
	os.Exit(0)
}
