
attribute max_length: Int64

attribute min_length: Int64

attribute fixed_length: Int64

//  (le <=)
attribute maximum: Int64

//  (le <=)
attribute exclusive_maximum: Int64

//  (ge >=)
attribute minimum: Int64

//  (ge >)
attribute exclusive_minmum: Int64

//  (ne !=)
attribute not_equal: Int64

//可以使用multiple_of关键字将数字限制为给定数字的倍数 。它可以设置为任何正数。
//{    "type": "number",    "multiple_of" : 10} 0 // OK10 // OK20 // OK23  // not OK，不是 10 的倍数
attribute multiple_of: UInt64
