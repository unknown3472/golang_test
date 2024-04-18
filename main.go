package main
import (
	"encoding/json"
	"log"
	"os"
)
type Employee struct{
	Id int `json:"id"`
	First_name string `json:"name"`
	Last_name string `json:"last_name"`
	Email string `json:"email"`
	Age int `json:"age"`
}
func main(){
	var e []Employee
	f, _ := os.ReadFile("employees.json")
	json.Unmarshal(f, &e)
	for j, i := range e{
		if i.Id==3{
			e[j].Age=50
		}
		if i.Id==6{
			e[j].Email, e[j].First_name, e[j].Last_name = "johndoe@gmail.com", "John", "Doe"
			e[j].Age= 100
		}
	}
	d, _ := json.MarshalIndent(e, "", " ")
	err := os.WriteFile("employees_new.json", d, 0644)
	if err!= nil{
		log.Fatal(err)
	}
}
