package cmd

import (
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/opsmanager"
	"github.com/spf13/viper"
)


func newClient() opsmanager.Client {
	baseURL := viper.GetString("baseURL")
	publicKey := viper.GetString("public_key")
	privateKey := viper.GetString("private_key")
	withResolver := opsmanager.WithResolver(httpclient.NewURLResolverWithPrefix(baseURL, opsmanager.PublicAPIPrefix))
	withDigestAuth := httpclient.WithDigestAuthentication(publicKey, privateKey)
	withHTTPClient := opsmanager.WithHTTPClient(httpclient.NewClient(withDigestAuth))

	return opsmanager.NewClient(withResolver, withHTTPClient)
}
