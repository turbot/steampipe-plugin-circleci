package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Handle rate limit errors
func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		// Check if any of the retry errors match the error message
		for _, retryError := range retryErrors {
			if strings.Contains(err.Error(), retryError) {
				plugin.Logger(ctx).Debug("circleci_errors.shouldRetryError", "retry_error", err)
				return true
			}
		}
		return false
	}
}
