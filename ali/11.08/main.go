package main
import ("fmt"
		"encoding/json"
)

//func main() {
//    fmt.Println(~2)
//}


type People struct {
    Name string `json:"name"`
}

func main() {
    js := `{
        "name":"seekload"
    }`
	fmt.Println([]byte(js))
    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println(p)
}
