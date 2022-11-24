package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciProject() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_project",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listCircleciProjects,
		},

		Columns: []*plugin.Column{
			{
				Name:        "raw",
				Description: "Raw data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
			},
		},
	}
}

//// LIST FUNCTION

func listCircleciProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV1(ctx, d)
	if err != nil {
		logger.Error("circleci_project.listCircleciProjects", "connect_error", err)
		return nil, err
	}

	projects, err := client.ListProjects()
	if err != nil {
		logger.Error("circleci_project.listCircleciProjects", "list_projects_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, project := range projects {
		d.StreamListItem(ctx, project)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
