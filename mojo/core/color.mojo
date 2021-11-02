
@format('#{red:0x}{green:0x}{blue:0x}')
type Color {
    red:   UInt8 @1
    green: UInt8 @2
    blue:  UInt8 @3
    alpha: Float @4
}

//type RGB = Color

//type HLS {
//}

//const {
//    red = '#ff0000'
//    blue = '#ff0000'
//}

//
//Color(str:String)

//
//Color(r:UInt8, g:UInt8, b:UInt8)

///
//func rgb(r:UInt8, g:UInt8, b:UInt8, a) -> Color

///
//func hsl(hue: Degree, saturation: Percentage, lightness: Percentage, alpha:Real) -> Color

///
//func lighten(c:Color, n:Integer) -> Color    // `red.ligthen(10)`

///
//func saturate(c:Color,)

///
//func spin(c:Color,)

///
//func fade(c:Color,)

///
//func mix(c:Color,)

///
//func darken(c:Color,)
