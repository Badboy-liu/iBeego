package enums

type Symbol string

const (
	Exact     Symbol = "exact"
	IExact    Symbol = "iexact"
	Like      Symbol = "contains"
	ILike     Symbol = "icontains"
	In        Symbol = "in"
	Ge        Symbol = "gt"
	Gte       Symbol = "gte"
	Le        Symbol = "lt"
	Lte       Symbol = "lte"
	TailLike  Symbol = "startswith"
	ITailLike Symbol = "istartswith"
	HeadLike  Symbol = "endswith"
	iHeadLike Symbol = "iendswith"
	IsNil     Symbol = "isnull"
)

type Operation int

const (
	Query  Operation = 0
	Create Operation = 1
	Update Operation = 2
	Delete Operation = 3
)
