/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/database"
	extension "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/extension"
	grant "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/grant"
	role "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/role"
	providerconfig "github.com/alexinthesky/provider-postgresql/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		extension.Setup,
		grant.Setup,
		role.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
