package helpers
import ("strconv"
"fmt")

// StringToUint return Uint
func StringToUint(value string)uint{
			intValue, err := strconv.Atoi(value)
					if err != nil{
						fmt.Println("Erro ao converter string para int:", err)
						fmt.Println(err)
					}
						uintValue := uint(intValue)
						return uintValue
}