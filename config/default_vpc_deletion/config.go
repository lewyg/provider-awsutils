package repository

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("awsutils_default_vpc_deletion", func(r *config.Resource) {
		r.ShortGroup = "default_vpc_deletion"
	})
}
