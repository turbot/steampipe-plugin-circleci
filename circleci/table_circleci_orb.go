package circleci

import (
	"context"
	"strings"

	"github.com/CircleCI-Public/circleci-cli/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableCircleciOrb() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_orb",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listCircleciOrbs,
		},

		Columns: []*plugin.Column{
			{Name: "orb_base", Description: "", Type: proto.ColumnType_JSON},
			{Name: "commands", Description: "", Type: proto.ColumnType_JSON},
			{Name: "jobs", Description: "", Type: proto.ColumnType_JSON},
			{Name: "executors", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listCircleciOrbs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV2Sdk(ctx, d)
	if err != nil {
		logger.Error("circleci_orb.listCircleciOrbs", "connect_error", err)
		return nil, err
	}

	var retrieveUncertifiedOrbs = false

	orbs, err := api.ListOrbs(client, retrieveUncertifiedOrbs)
	if err != nil {
		logger.Error("circleci_orb.listCircleciOrbs", "list_orbs_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, orb := range orbs.Orbs {
		d.StreamListItem(ctx, orb)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
