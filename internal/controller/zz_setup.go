/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	privileges "github.com/alexinthesky/provider-postgresql/internal/controller/default/privileges"
	replicationslot "github.com/alexinthesky/provider-postgresql/internal/controller/physical/replicationslot"
	database "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/database"
	extension "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/extension"
	function "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/function"
	grant "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/grant"
	publication "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/publication"
	role "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/role"
	schema "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/schema"
	server "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/server"
	subscription "github.com/alexinthesky/provider-postgresql/internal/controller/postgresql/subscription"
	providerconfig "github.com/alexinthesky/provider-postgresql/internal/controller/providerconfig"
	slot "github.com/alexinthesky/provider-postgresql/internal/controller/replication/slot"
	mapping "github.com/alexinthesky/provider-postgresql/internal/controller/user/mapping"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		privileges.Setup,
		replicationslot.Setup,
		database.Setup,
		extension.Setup,
		function.Setup,
		grant.Setup,
		publication.Setup,
		role.Setup,
		schema.Setup,
		server.Setup,
		subscription.Setup,
		providerconfig.Setup,
		slot.Setup,
		mapping.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
