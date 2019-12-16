package apiserver

// APIserver ...
type APIserver struct{
	config *Config
}

// New ...
func New() *APIserver {
	return &APIserver{
		config:config,
	}		
}

// Start ...
func (s *APIserver) Start() error  {
	return nil
}