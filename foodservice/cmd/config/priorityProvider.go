package config

type PriorityProvider struct {
	providers []ValueProvider
}

func NewPriorityProvider(providers ...ValueProvider) *PriorityProvider {
	return &PriorityProvider{providers: providers}
}

func (p *PriorityProvider) Load(key string) any {
	for _, provider := range p.providers {
		if val, ok := provider.Load(key); ok && val != nil {
			return val
		}
	}
	return nil
}
