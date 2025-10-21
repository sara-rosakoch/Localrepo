package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    var cards =[]int{2,6,9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index>=0 && index<=(len(slice)-1){
        return slice[index]
    }else{
        return -1
    }	
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index>=0 && index<=(len(slice)-1){
        slice[index]=value
        return slice 
    }else{
        return append(slice,value)
    }
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var cards []int
    cards= append(cards,values...)
    if 1<=len(values){
    	return append (cards,slice...)
        }else{
        return slice
    }
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index>=0 && index<=(len(slice)-1){
        a:=slice[:index]
        b:=slice[index+1:]
        return append(a,b...)
    }else{
        return slice
    }
}
