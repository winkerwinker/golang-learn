package bridge

type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s Square) Draw() string {
	return "square"
	//panic("implement me")
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{color: color, square: square}
}

func (c ColorSquare) Draw() string {
	//panic("implement me")
	// 调用的过程中组装了内部的参数，更类似与 模版方法
	return c.square.Draw() + "with color " + c.color

}
