///
///
@format('{$[:4]:%x}{-$[4:6]}{-$[6:8]}{-$[8:10]{-$[10:]}}')
type Uuid : Bytes @fixed_length(16)
