package oidc

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/authelia/authelia/v4/internal/configuration/schema"
)

var exampleIssuerPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAvcMVMB2vEbqI6PlSNJ4HmUyMxBDJ5iY7FS+zDDAHOZBg9S3S\nKcAn1CZcnyL0VvJ7wcdhR6oTnOwR94eKvzUyJZ+GL2hTMm27dubEYsNdhoCl6N3X\nyEEohNfoxiiCYraVauX8X3M9jFzbEz9+pacaDbHB2syaJ1qFmMNR+HSu2jPzOo7M\nlqKIOgUzA0741MaYNt47AEVg4XU5ORLdolbAkItmYg1QbyFndg9H5IvwKkYaXTGE\nlgDBcPUC0yVjAC15Mguquq+jZeQay+6PSbHTD8PQMOkLjyChI2xEhVNbdCXe676R\ncMW2R/gjrcK23zmtmTWRfdC1iZLSlHO+bJj9vQIDAQABAoIBAEZvkP/JJOCJwqPn\nV3IcbmmilmV4bdi1vByDFgyiDyx4wOSA24+PubjvfFW9XcCgRPuKjDtTj/AhWBHv\nB7stfa2lZuNV7/u562mZArA+IAr62Zp0LdIxDV8x3T8gbjVB3HhPYbv0RJZDKTYd\nzV6jhfIrVu9mHpoY6ZnodhapCPYIyk/d49KBIHZuAc25CUjMXgTeaVtf0c996036\nUxW6ef33wAOJAvW0RCvbXAJfmBeEq2qQlkjTIlpYx71fhZWexHifi8Ouv3Zonc+1\n/P2Adq5uzYVBT92f9RKHg9QxxNzVrLjSMaxyvUtWQCAQfW0tFIRdqBGsHYsQrFtI\nF4yzv8ECgYEA7ntpyN9HD9Z9lYQzPCR73sFCLM+ID99aVij0wHuxK97bkSyyvkLd\n7MyTaym3lg1UEqWNWBCLvFULZx7F0Ah6qCzD4ymm3Bj/ADpWWPgljBI0AFml+HHs\nhcATmXUrj5QbLyhiP2gmJjajp1o/rgATx6ED66seSynD6JOH8wUhhZUCgYEAy7OA\n06PF8GfseNsTqlDjNF0K7lOqd21S0prdwrsJLiVzUlfMM25MLE0XLDUutCnRheeh\nIlcuDoBsVTxz6rkvFGD74N+pgXlN4CicsBq5ofK060PbqCQhSII3fmHobrZ9Cr75\nHmBjAxHx998SKaAAGbBbcYGUAp521i1pH5CEPYkCgYEAkUd1Zf0+2RMdZhwm6hh/\nrW+l1I6IoMK70YkZsLipccRNld7Y9LbfYwYtODcts6di9AkOVfueZJiaXbONZfIE\nZrb+jkAteh9wGL9xIrnohbABJcV3Kiaco84jInUSmGDtPokncOENfHIEuEpuSJ2b\nbx1TuhmAVuGWivR0+ULC7RECgYEAgS0cDRpWc9Xzh9Cl7+PLsXEvdWNpPsL9OsEq\n0Ep7z9+/+f/jZtoTRCS/BTHUpDvAuwHglT5j3p5iFMt5VuiIiovWLwynGYwrbnNS\nqfrIrYKUaH1n1oDS+oBZYLQGCe9/7EifAjxtjYzbvSyg//SPG7tSwfBCREbpZXj2\nqSWkNsECgYA/mCDzCTlrrWPuiepo6kTmN+4TnFA+hJI6NccDVQ+jvbqEdoJ4SW4L\nzqfZSZRFJMNpSgIqkQNRPJqMP0jQ5KRtJrjMWBnYxktwKz9fDg2R2MxdFgMF2LH2\nHEMMhFHlv8NDjVOXh1KwRoltNGVWYsSrD9wKU9GhRCEfmNCGrvBcEg==\n-----END RSA PRIVATE KEY-----"

func TestOpenIDConnectProvider_NewOpenIDConnectProvider_NotConfigured(t *testing.T) {
	provider, err := NewOpenIDConnectProvider(nil)

	assert.NoError(t, err)
	assert.Nil(t, provider.Fosite)
	assert.Nil(t, provider.Store)
}

func TestOpenIDConnectProvider_NewOpenIDConnectProvider_BadIssuerKey(t *testing.T) {
	_, err := NewOpenIDConnectProvider(&schema.OpenIDConnectConfiguration{
		IssuerPrivateKey: "BAD KEY",
	})

	assert.Error(t, err, "abc")
}

func TestOpenIDConnectProvider_NewOpenIDConnectProvider_GoodConfiguration(t *testing.T) {
	provider, err := NewOpenIDConnectProvider(&schema.OpenIDConnectConfiguration{
		IssuerPrivateKey: exampleIssuerPrivateKey,
		HMACSecret:       "asbdhaaskmdlkamdklasmdlkams",
		Clients: []schema.OpenIDConnectClientConfiguration{
			{
				ID:     "a-client",
				Secret: "a-client-secret",
				Policy: "one_factor",
				RedirectURIs: []string{
					"https://google.com",
				},
			},
			{
				ID:          "b-client",
				Description: "Normal Description",
				Secret:      "b-client-secret",
				Policy:      "two_factor",
				RedirectURIs: []string{
					"https://google.com",
				},
				Scopes: []string{
					"groups",
				},
				GrantTypes: []string{
					"refresh_token",
				},
				ResponseTypes: []string{
					"token",
					"code",
				},
			},
		},
	})

	assert.NotNil(t, provider)
	assert.NoError(t, err)
}
