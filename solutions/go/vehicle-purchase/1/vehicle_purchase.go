package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck"{
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var chosen string
	if option1 <= option2{
        chosen = option1
    }else{
        chosen = option2
    }
        return chosen + " is clearly the better choice."
          
}


// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64
    
    if int(age)<3{
        price = float64(originalPrice*80/100)
    }else if int(age) >=3 && int(age)<10{
        price = float64(originalPrice*70/100)
    }else {
        price = float64(originalPrice*50/100)
    }
    return price 
}
