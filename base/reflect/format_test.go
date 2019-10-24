package reflect

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d = time.Second

	fmt.Println(Any(x))
	fmt.Println(Any(d))

	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

var strangelove = Movie{
	Title:    "大话西游",
	Subtitle: "大圣娶亲",
	Year:     1995,
	Color:    true,
	Actor: map[string]string{
		"周星驰": "主角",
		"吴孟达": "主角",
		"朱茵":  "主角",
		"蔡少芬": "配角",
		"蓝洁瑛": "配角",
		"莫文蔚": "配角",
		"罗家英": "配角",
		"陆树铭": "友情出演",
		"李健仁": "友情出演",
	},
	Oscars: []string{"最佳搞笑", "最佳主角", "最佳剧情"},
	Sequel: nil,
}

func TestDisplay(t *testing.T) {

	Display("大话西游", strangelove)

	//Display("os.Stderr", os.Stderr)
	Display("os.Stderr", reflect.ValueOf(os.Stderr))
}

func TestDisplay2(t *testing.T) {
	var i interface{} = 3
	Display("i", i)

	Display("&i", &i)
}

type Cycle struct {
	Value int
	Tail  *Cycle
}

// 存在循环引用的情况，会首尾相连，输出无穷无尽
func TestDisplay3(t *testing.T) {
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
}

type Student struct {
	Name  string
	Age   int
	Grade []int
	elder map[string]string
}

var stu = Student{
	Name:  "Eric",
	Age:   29,
	Grade: []int{560, 560, 580},
	elder: map[string]string{
		"爸爸":  "VV",
		"妈妈":  "YY",
		"爷爷":  "LHX",
		"奶奶":  "LLX",
		"外公":  "CMX",
		"曾祖母": "CSF",
	},
}

func TestMarshal(t *testing.T) {
	bytearr, _ := Marshal(stu)
	fmt.Printf("%s\n", string(bytearr))
}
