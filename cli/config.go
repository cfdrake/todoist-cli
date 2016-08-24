package cli

// Interface for a config provider.
type Configurer interface {
	UserToken() string
}
