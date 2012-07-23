package gfx

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Rect.h>
import "C"

// sfFloatRect and sfIntRect are utility classes for
// manipulating rectangles.
// \brief Check if a point is inside a rectangle's area
//
// \param rect Rectangle to test
// \param x    X coordinate of the point to test
// \param y    Y coordinate of the point to test
//
// \return sfTrue if the point is inside
//
// sfBool sfFloatRect_contains(const sfFloatRect* rect, float x, float y);

type FloatRect struct {
	Cref *C.sfFloatRect
}

func NewFloatRect(left, top, width, height float32) FloatRect {
	fr := C.sfFloatRect{
		C.float(left),
		C.float(top),
		C.float(width),
		C.float(height),
	}
	return FloatRect{&fr}
}

type IntRect struct {
	Cref *C.sfIntRect
}

func NewIntRect(left, top, width, height int32) IntRect {
	fr := C.sfIntRect{
		C.int(left),
		C.int(top),
		C.int(width),
		C.int(height),
	}
	return IntRect{&fr}
}

func (self FloatRect) Contains(x, y float32) bool {
	return C.sfFloatRect_contains(self.Cref, C.float(x), C.float(y)) == 1
}


func (self IntRect) Contains(x, y int32) bool {
	return C.sfIntRect_contains(self.Cref, C.int(x), C.int(y)) == 1
}

// brief Check intersection between two rectangles
// param rect1        First rectangle to test
// param rect2        Second rectangle to test
// param intersection Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfFloatRect_intersects(const sfFloatRect* rect1, const sfFloatRect* rect2, sfFloatRect* // intersection);
func (self FloatRect) Intersects(rect FloatRect) (*FloatRect, bool) {
	intersect := new(C.sfFloatRect)
	b := C.sfFloatRect_intersects(self.Cref, rect.Cref, intersect) == 1
	return &FloatRect{intersect}, b
}

// brief Check intersection between two rectangles
// param rect1        First rectangle to test
// param rect2        Second rectangle to test
// param intersection Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfIntRect_intersects(const sfIntRect* rect1, const sfIntRect* rect2, sfIntRect* // intersection);
func (self IntRect) Intersects(rect IntRect) (*IntRect, bool) {
	intersect := new(C.sfIntRect)
	b := C.sfIntRect_intersects(self.Cref, rect.Cref, intersect) == 1
	return &IntRect{intersect}, b
}

