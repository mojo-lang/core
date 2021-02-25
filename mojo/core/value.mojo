
type Object : {String: Value}

type Values : [Value]

type Value = Values @1 | Object @2 | Bool @3 | Int64 @4 | Double @5 | String @7 
