package lasagna

// // TODO: define the 'OvenTime' constant

// // RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
// func RemainingOvenTime(actualMinutesInOven int) int {
// 	panic("RemainingOvenTime not implemented")
// }

// // PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
// func PreparationTime(numberOfLayers int) int {
// 	panic("PreparationTime not implemented")
// }

// // ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
// func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
// 	panic("ElapsedTime not implemented")
// }
import "fmt"

const OvenTime int =40
func RemainingOvenTime(actual int ) int {
    return OvenTime - actual
}

func PreparationTime(numberOfLayers int )int{
    return numberOfLayers*2
}

func ElapsedTime(numberOfLayers int, actualMinutesInOven int)int{
    return PreparationTime(numberOfLayers )+actualMinutesInOven
}

func main(){
    
    fmt.Println("remaining time in the oven =" , RemainingOvenTime(25))
    fmt.Println("preperation time = ",PreparationTime(4))
    fmt.Println("elapsed time = ",ElapsedTime(4,25) )
    
    
}