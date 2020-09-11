package main
import ( 
	"fmt"
    "errors"
	"math"
)
//introducing structs
type person struct{
	name string
	age int
}
func main(){
	p := person{name:"Rachel",age:21}
	fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println(p.name)

	//getting the sum of digits from a function
	result:=sum(100,500)
	fmt.Println(result)
	outcome, err := sqrt(25)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(outcome)
	}
	//using the pointers, & for increment purposes and the asteric for notation
	i := 7
	inc(&i)
	fmt.Println(i)
	//declaring variables
    x:=89
	y:=7
	var sum int = x+y
	//using if and else statements

	if x>90{
		fmt.Println("X is smaller than ninety")

	}else if x<7{
		fmt.Println("Its greater than")
	}else{
		fmt.Println("loop out")
	}
	//working with arrays
	a:=[]int{5,4,3,2,1}
	a=append(a,13)
	vertices:=make(map[string]int)
	vertices["square"]=2
	vertices["triangle"]=3
	vertices["dodecagon"]=12
	delete(vertices,"square")
	//using the for loop in arrays
	for i:=0; i<5; i++{
		fmt.Println(i)
	}
	arr:=[]string{"a","b","c"}
	for index, value :=range arr{
		fmt.Println("index:",index, "value:",value)
	}
	//mapping through the array
	m:=make(map[string]string)
	m["a"]="alpha"
	m["b"]="beta"
	//dictionary included
	for key, value :=range arr{
		fmt.Println("key:", key , "value:", value)
	}
	

	fmt.Println(vertices["triangle"])
	fmt.Println(vertices)
	fmt.Println(a)
	fmt.Println(sum)
	fmt.Println("Hello World! This is my first Go program\n")
	

}
func inc(x * int ){
	*x++
}

//introduction to functions
func sum(x int , y int) int{
	return x+y
}
//finding squareroots and importing the math function

func sqrt(x float64)(float64,error){
	if x<0 {
		return 0, errors.New("Undefined for negative members")
	}
	return math.Sqrt(x), nil

}

