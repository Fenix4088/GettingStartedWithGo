package structspractice

func Init() {
	title := getUserData("Enter Product name: ")
	desc := getUserData("Shortly describe your product: ")
	price := getUserData("Enter Product price in USD: ")

	createProduct(title, desc, price).getProductInfo()

}