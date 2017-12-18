/// # 
///
///
type Array<T>

///
/// Creates an array with a given length, filled with a default element
///
/// ```
/// repeat(0, 5) == [0,0,0,0,0]
/// repeat("cat", 3) == ["cat","cat","cat"]
/// "cat".repeat(3) == ["cat","cat","cat"]
/// ```
func repeat<T>(T, UInt) -> Array<T>

/// ## Basics

/// 
func empty<T>([T])

func length<T>([T])

func push<T>([T], T) -> [T]

func append<T>([T], T) -> [T]

func append<T>([T], [T]) -> [T]

/// ## Get and Set
func get<T>([T], Int) -> Maybe<T>

func set<T>([T], Int, T) -> [T]