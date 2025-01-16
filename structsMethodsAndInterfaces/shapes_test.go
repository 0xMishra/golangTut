package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10.0, Height: 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, wanted %.2f", got, want)
	}
}

func TestArea1(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g, wanted %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		checkArea(t, rect, 100.0)
	})

	t.Run("circles", func(t *testing.T) {
		circ := Circle{10.0}
		checkArea(t, circ, 314.1592653589793)
	})
}

// table driven tests
func TestArea2(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Width: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shape.Area() != tt.hasArea {
				t.Errorf("%#v got %g, wanted %g", tt.shape, tt.shape.Area(), tt.hasArea)
			}
		})
	}
}
