package entity

type Category string

const (
	Boots    = "boots"
	Sandals  = "sandals"
	Sneakers = "sneakers"
)

var categoryTypeStrings = []Category{
	Boots,
	Sandals,
	Sneakers,
}

func CategoryName(n int) Category {
	return categoryTypeStrings[n]
}

func IsValid(s string) bool {
	if Category(s) == Boots || Category(s) == Sandals || Category(s) == Sneakers || s == "" {
		return true
	}
	return false
}
