# Table: circleci_insights_workflow_run

Get Insights for your workflows. It can access data spanning up to 90 days.

## Examples

### Get average duration and deployment count of a project for each month

```sql
select
  project_slug,
  to_char(created_at, 'YYYY-MM') as year_month,
  avg(duration) as average_duration,
  count(id) as deployment_count
from
  circleci_insights_workflow_run
where
  workflow_name = 'default'
  and project_slug = 'gh/companyname/projectname'
  and branch = 'main'
  and status = 'success'
group by
  project_slug,
  year_month;
```

### List workflows created in the last 30 days

```sql
select
  *
from
  circleci_insights_workflow_run
where
  workflow_name = 'default'
  and project_slug = 'gh/companyname/projectname'
  and created_at > current_date - interval '30' day;
```

### List workflows which are in failed state

```sql
select
  *
from
  circleci_insights_workflow_run 
where
  workflow_name = 'default'
  and project_slug = 'gh/companyname/projectname'
  and status = 'failed';
```
