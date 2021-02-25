
@template('#{red:0x}{green:}{blue}')
type Color {
    red: Int8
    green: Int8
    blue: Int8
    alpha: float
}

//type RGB = Color

//type HLS {
//}

const {
    red = '#ff0000'
    blue = '#ff0000'
}

///
Color(str:String)

///
Color(r:UInt8, g:UInt8, b:UInt8)

///
func rgb(r:UInt8, g:UInt8, b:UInt8, a) -> Color

///
func hsl(hue: Degree, saturation: Percentage, lightness: Percentage, alpha:Real) -> Color

///
func lighten(c:Color, n:Integer) -> Color    // `red.ligthen(10)`

///
func saturate(c:Color,)

///
func spin(c:Color,)

///
func fade(c:Color,)

///
func mix(c:Color,)

///
func darken(c:Color,)
