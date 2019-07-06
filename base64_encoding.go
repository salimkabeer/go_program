/*
we can encode String and url in Go. Go has Encoder which takes byte array
and convert into string encoding.

The Decoder takes the encoded value and covert it to the original string.

*/

package main  
  
import "fmt"  
import b64 "encoding/base64" 

func main() {  
    data := "JavaTpoint@12345!@#$%^&*()"  
    strEncode := b64.StdEncoding.EncodeToString([]byte(data))  
    fmt.Println("value to be encode  "+data)  
    fmt.Println("Encoden value:  "+strEncode)  
  
    fmt.Println()  
  
  
    fmt.Println("Value to be decode  "+strEncode)  
    strDecode, _ := b64.StdEncoding.DecodeString(strEncode)  
    fmt.Println("Decoded value  "+string( strDecode))  
    fmt.Println()  
  
    url := "https://golang.org/ref/spec"  
  
    fmt.Println("url to be encode  "+url)  
    urlEncode := b64.URLEncoding.EncodeToString([]byte(url))  
    fmt.Println("Encoded url  "+urlEncode)  
  
    fmt.Println("value to be decode  "+urlEncode)  
    strDecode2,_ := b64.URLEncoding.DecodeString(urlEncode)  
  
    fmt.Println("Decoded value  "+string(strDecode2))  
} 

/*
value to be encode  JavaTpoint@12345!@#$%^&*()
Encoden value:  SmF2YVRwb2ludEAxMjM0NSFAIyQlXiYqKCk=

Value to be decode  SmF2YVRwb2ludEAxMjM0NSFAIyQlXiYqKCk=
Decoded value  JavaTpoint@12345!@#$%^&*()

url to be encode  https://golang.org/ref/spec
Encoded url  aHR0cHM6Ly9nb2xhbmcub3JnL3JlZi9zcGVj
value to be decode  aHR0cHM6Ly9nb2xhbmcub3JnL3JlZi9zcGVj
Decoded value  https://golang.org/ref/spec
*/
