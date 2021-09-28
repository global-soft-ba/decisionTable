package data

type ExpressionInterface interface {
	String() string
	Convert(targetFormat OutputFormat) (string, error)
}
