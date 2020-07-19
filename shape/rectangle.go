package shape


type Rectangle struct {
	width float64
	height float64
}

/*C++式接口写法*/
func (r* Rectangle) SetWidth(w float64) {
	r.width = w
}

func (r* Rectangle) SetHeight(h float64) {
	r.height = h
}

func (r Rectangle) GetWidth() float64 {
	return r.width
}

func (r Rectangle) GetHeight() float64 {
	return r.height
}

/*Go式接口写法*/
func (r* Rectangle) SetRectangle(w float64, h float64) {
	r.height = h
    r.width = w
}
func (r Rectangle) GetRectangle() (float64, float64){
	return r.width, r.height
}

/*获取面积接口*/
func (r Rectangle) GetArea() (float64){
	return r.width * r.height
}

/*
实现非类成员函数，该函数接收两个矩形的实例，返回矩形面积的差值
*/
/*
func CalculateAreaDiff(a Rectangle, b Rectangle) float64 {
    return math.Abs(a.GetArea()-b.GetArea())
}
*/