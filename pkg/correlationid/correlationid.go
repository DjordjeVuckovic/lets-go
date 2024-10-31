package correlationid

type Id string
type contextKey string

const correlationKey contextKey = "correlation_id"

type Provider interface {
}
type Generator interface {
	Gen()
}
type Middleware interface {
}
