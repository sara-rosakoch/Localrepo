
//Package weather provide tools about current weather and past data.
package weather

var (
    //CurrentCondition represents current weather conditions.
	CurrentCondition string
    //CurrentLocation represents current location.
	CurrentLocation  string
)
//Forecast returns the weather forcast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
