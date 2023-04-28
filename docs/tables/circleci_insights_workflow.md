# Table: circleci_insights_workflow

Get Insights for your workflows

## Examples

### Get average duration and deployment count of a project for each month

```sql
SELECT 
    w.project_slug, to_char(w.created_at, 'YYYY-MM') as year_month, 
    AVG(w.duration) as average_duration, COUNT(w.id) as deployment_count
FROM 
    circleci_insights_workflow w
WHERE 
    w.workflow_name ='default' AND w.project_slug='gh/companyname/projectname' and w.branch ='main' AND w.status='success'
GROUP BY
    w.project_slug, year_month;
```
