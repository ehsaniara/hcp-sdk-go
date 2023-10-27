package tokencache

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
)

const (
	// refreshTimeout is the duration waited to receive a refresh token, before a new token is fetched.
	// If the refresh takes longer than 5 seconds it is probably quicker to just fetch a new token (if possible).
	refreshTimeout = 5 * time.Second

	// minTTL is the minimum time that a cached token has to be valid for in order to be returned. 15 seconds should be
	// sufficient for most calls that make use of a returned token.
	minTTL = 15 * time.Second
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

	// Return the access token if it is still valid for at least minTTL
	if token != nil && token.Expiry.After(time.Now().Add(minTTL)) {
		return token, nil
	}

	// Try to refresh the token if it has a RefreshToken and an oauth config was provided
	if token != nil && token.RefreshToken != "" && source.oauthConfig != nil {
		ctx, cancel := context.WithTimeout(context.Background(), refreshTimeout)
		defer cancel()

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
