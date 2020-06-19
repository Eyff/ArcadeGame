package invad

import(
)

//A Point2D is an object in 2d Space and holds an X and an Y coordinate
type Point2D struct{
	X,Y float64
}

//Creates a new Point2D 
func NewPoint(x,y float64) *Point2D{
	return &Point2D{x,y}
}

//A Rectangle stores the 2dim information on its four sides 
type Rectangle2D struct{
	L,U,R,D float64
}

//New Rectangle2D 
func NewRectangle2D(l,u,r,d float64) *Rectangle2D{
	return &Rectangle2D{l,u,r,d}
}

//Within checks if a pair of x,y coordinates are within a Rectangle2D
func (r2d Rectangle2D)Within(x,y float64) bool{
	if x > r2d.L && y > r2d.L && x < r2d.R && y < r2d.D{
		return true
	}
	return false
}

//PointWithin checks if a Point2D is wihtin a Rectangle2D
func (r2d Rectangle2D)PointWithin(p2d *Point2D) bool{
	if p2d.X > r2d.L && p2d.Y > r2d.U && p2d.X < r2d.R && p2d.Y < r2d.D{
		return true
	}
	return false
}

//Intersects checks if two Rectangle intersect each other
func (r2d Rectangle2D)Intersects(R2D *Rectangle2D)bool{
	if r2d.PointWithin(NewPoint(R2D.L, R2D.U)) || r2d.PointWithin(NewPoint(R2D.L, R2D.D))  || r2d.PointWithin(NewPoint(R2D.R,R2D.U)) ||  r2d.PointWithin(NewPoint(R2D.R, R2D.D)){
		return true
	}
	if R2D.PointWithin(NewPoint(r2d.L, r2d.U)) || R2D.PointWithin(NewPoint(r2d.L, r2d.D)) || R2D.PointWithin(NewPoint(r2d.R, r2d.U)) || R2D.PointWithin(NewPoint(r2d.R, r2d.D)){
		return true
	}
	return false
}

//Move moves the coordinates of a Rectangle2D
func (r2d *Rectangle2D)Move(x,y float64){
	r2d.L += x
	r2d.R += x
	r2d.U += y
	r2d.D += y
}

//Places the left up corner at a given Point2D coordinates
func (r2d *Rectangle2D)Place(p2d *Point2D){
	length := r2d.R - r2d.L
	height := r2d.D - r2d.U
	r2d.L = p2d.X
	r2d.R = p2d.X + length
	r2d.U = p2d.Y
	r2d.D = p2d.Y + height
} 
