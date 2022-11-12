package option

type ApplyingOption interface {
	Apply(o *Option)
}
