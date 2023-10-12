package main
import "fmt"
func main(){
	i:=1
	for i<=3{
		fmt.Println(i)
		i++
		//i=i+1
	}
	fmt.Println("loop end ,new one coming")
	for j:=7;j<=10;j++{
		fmt.Println(j)
	}
}