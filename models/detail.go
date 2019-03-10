package models
type photo struct {
	Id string
	ImagePath string
}

//返回数据
func Getphoto() []photo {
	buffer :=[]photo{
		photo{
			Id:"pic1",
			ImagePath:"img01.png",
		},
		photo{
		Id:"pic2",
		ImagePath:"img02.png",
		},
		photo{
		Id:"pic3",
		ImagePath:"img03.png",
		},
		photo{
		Id:"pic4",
		ImagePath:"img04.png",
		},
		photo{
		Id:"pic5",
		ImagePath:"img05.png",
		},
		photo{
		Id:"pic6",
		ImagePath:"img06.png",
		},
	}
	return buffer
}

func GetTitle() string{
	return "fads 发生地方"
}