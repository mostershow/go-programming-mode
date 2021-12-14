package __common_method

import "fmt"

// 把一个结构体给嵌到另一个结构体 把 Widget嵌入到 Label 中
type Widget struct {
	X, Y int
}
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

// 层层嵌套，Label再嵌入到Button中
type Button struct {
	Label // Embedding (delegation)
}

func NewButton(i int, i2 int, s string) Button {
	return Button{Label{Widget{i, i2}, s}}
}

// Widget嵌入到ListBox中
type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

//////////////////////////////////////////////
// 行为1 展示
type Painter interface {
	Paint()
}

// 行为2 点击
type Clicker interface {
	Click()
}

/////////////////////////////////////////////

// 下面是一些实现
func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中可以重载这个接口方法以
func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}
func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

//////////////////////////////////////

func main() {
	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}
	for _, painter := range []Painter{listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}
