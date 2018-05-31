package rpconfig

type rpcConfig struct {
	serName string
	host    string
}
type ConfOption func(c *rpcConfig)

func WithSerName(s string) ConfOption {
	return func(opts *rpcConfig) {
		opts.serName = s
	}
}

func Init(c ...ConfOption) {
	opts := &rpcConfig{}

	for _, f := range c {
		f(opts)
	}
}
