package main

import (
	"encoding/json"
	"fmt"
)

//Json marshaling

// type Employee struct{
// 	Name string 
// 	Age int
// 	Salary float64
// }

func Marshalling(){
	//Json marshaling
	emp := Employee{
		Name: "Ganga",Age: 22,Salary: 25000,
	}
	fmt.Println(emp)

	jsonData,err :=json.Marshal(emp)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println(jsonData)
	fmt.Println(string(jsonData))

}

//Json Unmarshalling
type Employee struct{
	Name string
	Age int
	Salary float64 
}


func Unmarshalling(){

	var jsonString = `{"Name":"Ganga","Age":22,"Salary":25000}`

	var emp Employee

	error :=json.Unmarshal([]byte(jsonString),&emp)
	if error != nil{
		fmt.Println(error)
		return
	}
	fmt.Println(jsonString)
	fmt.Println(emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)  
	fmt.Println(emp.Salary)


}

func ArrayUnmarashaling(){

	type Location struct{
		City string
		Pincode int
	}


	var data = `[
	{"City":"Pune","Pincode":43211},
	{"City":"Mumbai","Pincode":43781},
	{"City":"Parbhani","P1incode":431402}]`

	var location []Location

	error :=json.Unmarshal([]byte(data),&location)
	if error != nil{
		fmt.Println(error)
		return
	}
	for _,val := range location{
		fmt.Println(val.City,val.Pincode)

	}
}

func CustomejsonUnMarshalling(){
	var data = `[
	{"City":"Pune","Pincode":43211},
	{"City":"Mumbai","Pincode":43781},
	{"City":"Parbhani","Pincode":431402}]`

	var decode []map[string]interface{}

	error := json.Unmarshal([]byte(data),&decode)
	if error!=nil{
		fmt.Println(error)
		return
	}
	fmt.Println(decode)
	for _,val := range decode{fmt.Println(val["City"],val["Pincode"])


	}

}

func Customemarshaling(){

	var data1 map[string]interface{} = make(map[string]interface{})

	data1["Name"] = "raju"
	data1["Age"] = 22
	data1["Pincode"] = 431402

	jsonData,error := json.Marshal(data1)
	if error != nil{
		fmt.Println(error)
		return
	}
	fmt.Println(string(jsonData))
}

func Embeddedmarshalling(){
	type Human struct{
		Name string
		Age int

	}
	type Man struct{
		Weight float64
		Height float32
		Human
	}
	var man Man
	man.Name = "Raju"
	man.Age = 25
	man.Height = 162
	man.Weight = 60.2
	
	jdata,err := json.Marshal(man)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(man)
	fmt.Println(string(jdata))
}

 
func main(){
	//Marshalling()
	//Unmarshalling()
	//ArrayUnmarashaling()
	//CustomejsonUnMarshalling()
	//Customemarshaling()
	Embeddedmarshalling()
	

	



}