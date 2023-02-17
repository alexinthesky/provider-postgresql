/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/base64"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/alexinthesky/provider-postgresql/apis/v1beta1"
	"github.com/upbound/upjet/pkg/terraform"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal postgresql credentials as JSON"
	// config keys
	keyHost           = "host"
	keyPort           = "port"
	keyDatabase       = "database"
	keyUsername       = "username"
	keyPassword       = "password"
	keySslMode        = "sslmode"
	keyConnectTimeout = "connect_timeout"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		ps.Configuration = map[string]any{}
		hostSecret := pc.Spec.Credentials.CommonCredentialSelectors
		hostSecret.SecretRef.Key = keyHost
		hostValueb64, err := resource.ExtractSecret(ctx, client, hostSecret)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		hostValue, err := base64.StdEncoding.DecodeString(string(hostValueb64))
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials+string(hostValueb64))
		}
		ps.Configuration[keyHost] = hostValue

		portSecret := pc.Spec.Credentials.CommonCredentialSelectors
		portSecret.SecretRef.Key = keyPort
		portValueb64, err := resource.ExtractSecret(ctx, client, portSecret)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		portValue, err := base64.StdEncoding.DecodeString(string(portValueb64))
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials+string(portValueb64))
		}
		ps.Configuration[keyPort] = string(portValue)

		usernameSecret := pc.Spec.Credentials.CommonCredentialSelectors
		usernameSecret.SecretRef.Key = keyUsername
		usernameValueb64, err := resource.ExtractSecret(ctx, client, usernameSecret)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		usernameValue, err := base64.StdEncoding.DecodeString(string(usernameValueb64))
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		ps.Configuration[keyUsername] = usernameValue

		passwordSecret := pc.Spec.Credentials.CommonCredentialSelectors
		passwordSecret.SecretRef.Key = keyPassword
		passwordValueb64, err := resource.ExtractSecret(ctx, client, passwordSecret)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		passwordValue, err := base64.StdEncoding.DecodeString(string(passwordValueb64))
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		ps.Configuration[keyPassword] = passwordValue

		// creds := map[string]string{}
		// if err := json.Unmarshal(data, &creds); err != nil {
		// 	return ps, errors.Wrap(err, errUnmarshalCredentials)
		// }

		// // Set credentials in Terraform provider configuration.
		// ps.Configuration = map[string]any{}
		// if v, ok := creds[keyHost]; ok {
		// 	ps.Configuration["keyHost] = v
		// }
		// if v, ok := creds[keyPort]; ok {
		// 	ps.Configuration["keyPort"] = v
		// }
		// if v, ok := creds[keyDatabase]; ok {
		// 	ps.Configuration["keyDatabase"] = v
		// }
		// if v, ok := creds[keyUsername]; ok {
		// 	ps.Configuration["keyUsername"] = v
		// }
		// if v, ok := creds[keyPassword]; ok {
		// 	ps.Configuration["keyPassword"] = v
		// }
		// if v, ok := creds[keySslMode]; ok {
		// 	ps.Configuration["keySslMode"] = v
		// }
		// if v, ok := creds[keyConnectTimeout]; ok {
		// 	ps.Configuration["keyConnectTimeout"] = v
		// }
		return ps, nil
	}
}
