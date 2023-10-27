package tokencache

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
)

// sourceType identities the type of token source.
type sourceType = string

// cachingTokenSource acts as a read-through cache for token information received from token sources and oauth configurations.
type cachingTokenSource struct {
	cacheFile        string
	login            bool
	forceLogin       bool
	sourceType       sourceType
	sourceIdentifier string
	oauthTokenSource oauth2.TokenSource
	oauthConfig      oAuth2Config
	fetchTimeout     time.Duration
}

// Token implements the oauth2.TokenSource interface. It will read cached tokens from a file and based on their validity
// return, refresh or replace them.
func (source *cachingTokenSource) Token() (*oauth2.Token, error) {
	// Read the cache information from the file, if it exists
	cachedTokens, err := readCache(source.cacheFile)
	if err != nil {
		log.Println(err)
	}

	// Garbage collect expired tokens
	cachedTokens.removeExpiredTokens()

	// Handle login tokens
	if source.login {
		return source.loginToken(cachedTokens)
	}

	// Handle non-login tokens
	switch source.sourceType {
	case sourceTypeServicePrincipal:
		return source.servicePrincipalToken(cachedTokens)
	case sourceTypeWorkload:
		return source.workloadToken(cachedTokens)
	default:
		return nil, fmt.Errorf("invalid source type: %q", source.sourceType)
	}
}

// getValidToken will perform the following steps:
// 1. check if the cached token is still valid and return it if this is the case
// 2. try to refresh the cached token using the provided oauth config
// 3. fetch a new token from the provided token source
func (source *cachingTokenSource) getValidToken(hitEntry *cacheEntry) (*oauth2.Token, error) {
	var token *oauth2.Token
	if hitEntry != nil {
		token = hitEntry.token()
	}

	// Return the access token if it is still valid
	if token != nil && token.Expiry.After(time.Now()) {
		return token, nil
	}

	// Try to refresh the token if it has a RefreshToken and an oauth config was provided
	if token != nil && token.RefreshToken != "" && source.oauthConfig != nil {
		ctx := context.Background()
		//ctx, cancel := context.WithTimeout(context.Background(), source.fetchTimeout)
		//defer cancel()

		token, err := source.oauthConfig.TokenSource(ctx, token).Token()
		if err == nil {
			return token, err
		}

		// Fall through to fetch a new token
		log.Printf("failed to refresh the token: %s\n", err)
	}

	// Fetch a new token
	if source.oauthTokenSource != nil {
		// Try to get a new token
		return source.oauthTokenSource.Token()
	}

	return nil, fmt.Errorf("no valid credential source available")
}
