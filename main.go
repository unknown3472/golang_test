package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
type User struct {
    Name       string
    Age        int
    Occupation string
}

func main() {
    file, err := os.Open("sample.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var users []User
    var name, occupation string
    var age int
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            if name != "" {
                users = append(users, User{Name: name, Age: age, Occupation: occupation})
                name, occupation = "", ""
                age = 0
            }
        } else {
            parts := strings.Split(line, ":")
            if len(parts) == 2 {
                key := strings.TrimSpace(parts[0])
                value := strings.TrimSpace(parts[1])

                switch key {
                case "Name":
                    name = value
                case "Age":
                    fmt.Sscanf(value, "%d", &age)
                case "Occupation":
                    occupation = value
                }
            }
        }
    }
    if name != "" {
        users = append(users, User{Name: name, Age: age, Occupation: occupation})
    }
    for _, user := range users {
        fmt.Printf("Name: %s, Age: %d, Occupation: %s\n", user.Name, user.Age, user.Occupation)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
    }
}


