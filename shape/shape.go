package shape

/*
将圆形和矩形共享接口独立出来
类似于c++的基类，函数为纯虚函数
*/
type Shape interface {
    GetArea()
}