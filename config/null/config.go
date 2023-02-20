/*
Copyright 2021 Upbound Inc.
*/

package postgresql

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("postgresql_database", func(r *ujconfig.Resource) {
		r.Kind = "Resource"
		// And other overrides.
	})
	p.AddResourceConfigurator("postgresql_role", func(r *ujconfig.Resource) {
		r.Kind = "Resource"
		// And other overrides.
	})
}
