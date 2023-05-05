package visitor

import "math"

type Shape interface {
	Accept(visitor CalculateVisitor)
}

type CalculateVisitor interface {
	VisitSquare(s *Square)
	VisitRectangle(r *Rectangle)
	VisitCircle(c *Circle)
}

// 正方形
type Square struct {
	SideLength float64
}

func (s *Square) Accept(visitor CalculateVisitor) {
	visitor.VisitSquare(s)
}

// 长方形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor CalculateVisitor) {
	visitor.VisitRectangle(r)
}

// 圆形
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor CalculateVisitor) {
	visitor.VisitCircle(c)
}

// 面积访问者
type AreaVisitor struct {
	Area float64
}

func (av *AreaVisitor) VisitSquare(s *Square) {
	av.Area = s.SideLength * s.SideLength
}

func (av *AreaVisitor) VisitRectangle(r *Rectangle) {
	av.Area = r.Width * r.Height
}

func (av *AreaVisitor) VisitCircle(c *Circle) {
	av.Area = math.Pi * c.Radius * c.Radius
}

// 周长访问者
type PerimeterVisitor struct {
	Perimeter float64
}

func (pv *PerimeterVisitor) VisitSquare(s *Square) {
	pv.Perimeter = 4 * s.SideLength
}

func (pv *PerimeterVisitor) VisitRectangle(r *Rectangle) {
	pv.Perimeter = 2 * (r.Width + r.Height)
}

func (pv *PerimeterVisitor) VisitCircle(c *Circle) {
	pv.Perimeter = 2 * math.Pi * c.Radius
}
