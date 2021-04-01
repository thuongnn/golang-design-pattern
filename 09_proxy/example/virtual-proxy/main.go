package main

import "fmt"

/*------------------- explain -------------------*/
/*
### Lợi ích của Proxy Pattern là gì?
- Cãi thiện Performance thông qua lazy loading, chỉ tải các tài nguyên khi chúng được yêu cầu.
- Nó cung cấp sự bảo vệ cho đối tượng thực từ thế giới bên ngoài.
- Giảm chi phí khi có nhiều truy cập vào đối tượng có chi phí khởi tạo ban đầu lớn.
- Dễ nâng cấp, bảo trì.

### Sử dụng Proxy Pattern khi nào?
- Khi muốn bảo vệ quyền truy xuất vào các phương thức của object thực.
- Khi cần một số thao tác bổ sung trước khi thực hiện phương thức của object thực.
- Khi tạo đối tượng ban đầu là theo yêu cầu hoặc hệ thống yêu cầu sự chậm trễ khi tải một số tài nguyên nhất định (lazy loading).
- Khi có nhiều truy cập vào đối tượng có chi phí khởi tạo ban đầu lớn.
- Khi đối tượng gốc tồn tại trong môi trường từ xa (remote).
- Khi đối tượng gốc nằm trong một hệ thống cũ hoặc thư viện của bên thứ ba.
- Khi muốn theo dõi trạng thái và vòng đời đối tượng.
*/
/*-----------------------------------------------*/


/*-- Có nhiều loại proxy khác nhau, dưới đây là vd của virtual proxy --*/
type Image interface {
	showImage()
}

/*------------------- setup real image -------------------*/
type RealImage struct {
	url string
}

func NewRealImage(url string) *RealImage {
	fmt.Printf("Image loaded: %s\n", url)
	return &RealImage{url}
}

func (ri *RealImage) showImage() {
	fmt.Printf("Image showed: %s\n", ri.url)
}

/*------------------- setup proxy image -------------------*/
type ProxyImage struct {
	realImage *RealImage
	url       string
}

func NewProxyImage(url string) *ProxyImage {
	fmt.Printf("Image unloaded: %s\n", url)
	return &ProxyImage{url: url}
}

func (pi *ProxyImage) showImage() {
	if pi.realImage == nil {
		pi.realImage = NewRealImage(pi.url)
	} else {
		fmt.Printf("Image already existed: %s\n", pi.url)
	}

	pi.realImage.showImage()
}

/*------------------- running -------------------*/
func main() {
	var proxyImage Image

	fmt.Println("Init proxy image: ")
	proxyImage = NewProxyImage("https://gpcoder.com/favicon.ico")

	fmt.Println("---")
	fmt.Println("Call real service 1st: ")
	proxyImage.showImage()

	fmt.Println("---")
	fmt.Println("Call real service 2st: ")
	proxyImage.showImage()
}


