/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"postgresql_database":                  config.NameAsIdentifier,
	"postgresql_default_privileges":        config.IdentifierFromProvider,
	"postgresql_extension":                 config.IdentifierFromProvider,
	"postgresql_function":                  config.IdentifierFromProvider,
	"postgresql_grant":                     config.NameAsIdentifier,
	"postgresql_physical_replication_slot": config.NameAsIdentifier,
	"postgresql_publication":               config.NameAsIdentifier,
	"postgresql_replication_slot":          config.NameAsIdentifier,
	"postgresql_role":                      config.NameAsIdentifier,
	"postgresql_schema":                    config.NameAsIdentifier,
	"postgresql_server":                    config.IdentifierFromProvider,
	"postgresql_subscription":              config.NameAsIdentifier,
	"postgresql_user_mapping":              config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
