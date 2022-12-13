package gifdraw

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black} //存放颜色
const whiteIndex = 0
const blackIndex = 1

//画一个Lissajous曲线的GIF图像，然后按照GIF格式进行编码，并写入到输出设备中。
func Lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的x轴Lissajous图形生成器的周期数
		res     = 0.001 //角坐标分辨率
		size    = 100   //图像画布所覆盖的大小[-size..+size]
		nframes = 64    //动画的桢数
		delay   = 8     //每桢之间以10毫秒为间隔的延迟
	)
	freq := rand.Float64() * 3.0 //与y轴Lissajous图形生成器的相对频率

	anim := gif.GIF{LoopCount: nframes} //声明并构造（通过{ } 语法）了一个GIF图形
	phase := 0.0                        // 相差（phase differences）
	//循环，画出每一桢的图形
	for i := 0; i < nframes; i++ {
		//创建了一个抽象的几何矩形对象（0，0）-（201，201）
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		//创建了一个画板对象， 画板对象是在几何图形的基础上增加了颜色，默认颜色是paltte的第0个元素，也就是白色。
		//img是一个指针
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//设置图形上点的色素，在白色图形上画出黑色的点，形成黑色的线。
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		//移相，以便画出另外一个相差不同的正弦曲线。
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		//动画anim的Image属性是一个指针切片（slice），相当于数组。
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
