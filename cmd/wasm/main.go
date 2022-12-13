//在文件中规定上述编译条件，引入"syscall/js"包时vscode才不会报错。
//否则就要在.vscode配置文件的settings.json中设置项目全局的编译环境为:
/**
{
  "go.toolsEnvVars": {
    "GOOS": "js",
    "GOARCH": "wasm"
  }
}
**/
//但项目全局的编译环境设定会影响其他目标环境不是js与wasm下的代码的vscode 测试。

package main

import (
	// 这是一个处于试验状态的包，Go团队不保证其向后兼容。
	"bytes"
	"encoding/base64"
	"fmt"
	"syscall/js"

	gifdraw "example.com/webassembly/cmd/wasm/gif"
	jsonhelper "example.com/webassembly/cmd/wasm/json"
	"example.com/webassembly/cmd/wasm/qrcode"
)

// go语言与JS语言进行互操作的库主要是"syscall/js"，但是这个库还处于试验阶段，
//不保证API的兼容性。
func main() {

	//使得jsonFormater函数成为全局的JS函数，并命名为formatJSON。这样就可以在浏览器的控制台中调用。
	//向window对象注册了个一个名为formatJSON的全局Js函数。
	js.Global().Set("formatJSON", js.FuncOf(jsonFormater))
	//向window对象注册了个一个名为drawQR的全局Js函数。
	js.Global().Set("drawQR", js.FuncOf(qrGenerator))
	//向window对象注册了个一个名为operateDom的全局Js函数。
	js.Global().Set("operateDom", js.FuncOf(operateDom))
	//向window对象注册了个一个名为drawGif的全局Js函数。
	js.Global().Set("drawGif", js.FuncOf(drawGif))

	fmt.Println("Go Web Assembly")
	//为了防止GO程序在浏览器中运行后退出，导致webassembly代码失效，增加下面一行，让goroutine保持等待。
	<-make(chan struct{})

}

//这个函数调用gifdraw包中的Lissajous函数画出动画，并转换为base64代码，以便在HTML img元素中展示。
var drawGif = func(this js.Value, args []js.Value) interface{} {
	var buf bytes.Buffer
	gifdraw.Lissajous(&buf)
	gifStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	return gifStr

}

//这个函数展示了如何在GO中操作浏览器的DOM对象。包括得到window对象，document对象等。
var operateDom = func(this js.Value, args []js.Value) interface{} {
	window := js.Global()         //window代表网页（page）的根（root）对象（统一用js.Valuel类型代表js对象），可以根访问页面的任何元素。
	doc := window.Get("document") // js.Value对象的Get方法返回给定属性名称的另一个js.Value类型的js对象。
	body := doc.Get("body")
	//用document对象的createElement方法创建一个div元素（后面会将该div元素挂接到document的body上）
	div := doc.Call("createElement", "div") //调用js.Value所代表的js对象的方法，第一个参数时当法名，后面跟着参数列表（一个或多个参数）。
	div.Set("textContent", "hello!!")       //设置js对象的属性，与Get方法相对应。
	body.Call("appendChild", div)           //为body挂接先前创建的元素。
	//GO WASM统一格式的可封装为JS函数的GO函数，用做html文档body元素onclick事件的处理函数。
	var bodyClickHandler = func(this js.Value, args []js.Value) interface{} {
		div := doc.Call("createElement", "div")
		div.Set("textContent", "click!!")
		body.Call("appendChild", div)
		return nil
	}
	//设置Body的js.FuncOf返回一个js.Func对象（结构体），类似于js.Value代表Js对象，js.Func代表JS函数。
	body.Set("onclick", js.FuncOf(bodyClickHandler))
	return "ok"
}

// 这个函数将会被封装JS函数，用于生成二维码图片的base64位编码，以便在HTML img 元素中展现。
var qrGenerator = func(this js.Value, args []js.Value) interface{} {
	fmt.Println("generate qr")
	if len(args) == 0 {
		fmt.Println("error,no input contend!")
		return ""
	}
	inputContent := args[0].String()
	fmt.Println(inputContent)
	if base64Str, err := qrcode.GenerateQr(inputContent); err != nil {
		return err
	} else {
		return base64Str
	}
}

//凡在浏览器要被调用的函数，也就是要被包装为Js语言所调用函数的GO函数，都要写成如下形式
type FuncCallByJS = func(this js.Value, args []js.Value) interface{}

var jsonFormater = func(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid no of arguments passed"
	}
	inputJSON := args[0].String()
	fmt.Printf("input %s\n", inputJSON)
	pretty, err := jsonhelper.PrettyJson(inputJSON)
	if err != nil {
		fmt.Printf("unable to convert to json %s\n", err)
		return err.Error()
	}
	return pretty
}

func jsWrapper(fun FuncCallByJS) js.Func {
	return js.FuncOf(fun)
}
