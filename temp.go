package main
import "fmt"

func NewMapEntry() map[string]interface{} {
    return map[string]interface{}{
        "path": "/var/log/**/*.log",
        "regex": "some-regex",
    }
}

func main() {
    fmt.Println("Hello World")
    
    a := NewMapEntry()
    b := NewMapEntry()
    
    arr := []interface{}{a, b}
    
    for _, pathObj := range arr {
        if obj, ok := pathObj.(map[string]interface{}); ok {
            path, _ := obj["path"].(string)
            regex, _ := obj["regex"].(string)
            
            fmt.Println(path, regex)
        }
    }
}
