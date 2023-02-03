package static

var wishlist Wishlist

type Wishlist struct {
	Name     string   `json:"wishlist_name"`
	Wishlist []string `json:"wishlist"`
}

func CreateWishlist() {
	var new_Wishlist Wishlist
	wishlist = new_Wishlist
}

func GetWishlist() Wishlist {
	return wishlist
}

func SetWishlistName(name string) {
	wishlist.Name = name
}

func AddToWishlist(item string) {
	wishlist.Wishlist = append(wishlist.Wishlist, item)
}

func RemoveFromWishlist(item string) {
	//REMOVING FIRST PRODUCT WITH THE SAME NAME
	var index int
	for i := 0; i < len(wishlist.Wishlist); i++ {
		if item == wishlist.Wishlist[i] {
			index = i
			break
		}
	}
	wishlist.Wishlist = removeIndex(wishlist.Wishlist, index)
}

func removeIndex(wishlist []string, index int) []string {
	return append(wishlist[:index], wishlist[index+1:]...)
}

func EmptyWishlist() {
	wishlist.Wishlist = nil
}
