package metrics

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// QueryOptsBuilder allows extensions to add parameters to the Query request.
type QueryOptsBuilder interface {
	ToMetricQueryQuery() (string, error)
}

// QueryOpts contains the options for a Prometheus instant query.
type QueryOpts struct {
	// Query is the PromQL query string.
	Query string `q:"query" required:"true"`

	// Time is the evaluation timestamp (RFC3339 or Unix timestamp).
	Time string `q:"time"`

	// Timeout is the evaluation timeout.
	Timeout string `q:"timeout"`
}

// ToMetricQueryQuery formats QueryOpts into a query string.
func (opts QueryOpts) ToMetricQueryQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Query performs a Prometheus instant query.
func Query(ctx context.Context, client *gophercloud.ServiceClient, opts QueryOptsBuilder) (r QueryResult) {
	_ = "STUB: not implemented"
	return *new(QueryResult)
}

// LabelsOptsBuilder allows extensions to add parameters to the Labels request.
type LabelsOptsBuilder interface {
	ToMetricLabelsQuery() (string, error)
}

// LabelsOpts contains the options for listing label names.
type LabelsOpts struct {
	// Match is a list of series selectors to filter labels by.
	Match []string `q:"match[]"`

	// Start is the start timestamp (RFC3339 or Unix timestamp).
	Start string `q:"start"`

	// End is the end timestamp (RFC3339 or Unix timestamp).
	End string `q:"end"`
}

// ToMetricLabelsQuery formats LabelsOpts into a query string.
func (opts LabelsOpts) ToMetricLabelsQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Labels returns a list of label names.
func Labels(ctx context.Context, client *gophercloud.ServiceClient, opts LabelsOptsBuilder) (r LabelsResult) {
	_ = "STUB: not implemented"
	return *new(LabelsResult)
}

// LabelValuesOptsBuilder allows extensions to add parameters to the LabelValues request.
type LabelValuesOptsBuilder interface {
	ToMetricLabelValuesQuery() (string, error)
}

// LabelValuesOpts contains the options for listing label values.
type LabelValuesOpts struct {
	// Match is a list of series selectors to filter label values by.
	Match []string `q:"match[]"`

	// Start is the start timestamp (RFC3339 or Unix timestamp).
	Start string `q:"start"`

	// End is the end timestamp (RFC3339 or Unix timestamp).
	End string `q:"end"`
}

// ToMetricLabelValuesQuery formats LabelValuesOpts into a query string.
func (opts LabelValuesOpts) ToMetricLabelValuesQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LabelValues returns a list of label values for a given label name.
func LabelValues(ctx context.Context, client *gophercloud.ServiceClient, name string, opts LabelValuesOptsBuilder) (r LabelValuesResult) {
	_ = "STUB: not implemented"
	return *new(LabelValuesResult)
}

// SeriesOptsBuilder allows extensions to add parameters to the Series request.
type SeriesOptsBuilder interface {
	ToMetricSeriesQuery() (string, error)
}

// SeriesOpts contains the options for finding series by label matchers.
type SeriesOpts struct {
	// Match is a list of series selectors. At least one must be provided.
	Match []string `q:"match[]" required:"true"`

	// Start is the start timestamp (RFC3339 or Unix timestamp).
	Start string `q:"start"`

	// End is the end timestamp (RFC3339 or Unix timestamp).
	End string `q:"end"`
}

// ToMetricSeriesQuery formats SeriesOpts into a query string.
func (opts SeriesOpts) ToMetricSeriesQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Series returns the list of time series that match certain label sets.
func Series(ctx context.Context, client *gophercloud.ServiceClient, opts SeriesOptsBuilder) (r SeriesResult) {
	_ = "STUB: not implemented"
	return *new(SeriesResult)
}

// TargetsOptsBuilder allows extensions to add parameters to the Targets request.
type TargetsOptsBuilder interface {
	ToMetricTargetsQuery() (string, error)
}

// TargetsOpts contains the options for listing targets.
type TargetsOpts struct {
	// State filters targets by state ("active", "dropped", "any").
	State string `q:"state"`
}

// ToMetricTargetsQuery formats TargetsOpts into a query string.
func (opts TargetsOpts) ToMetricTargetsQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Targets returns an overview of the current state of Prometheus target discovery.
func Targets(ctx context.Context, client *gophercloud.ServiceClient, opts TargetsOptsBuilder) (r TargetsResult) {
	_ = "STUB: not implemented"
	return *new(TargetsResult)
}

// RuntimeInfo returns runtime information about the Prometheus server.
func RuntimeInfo(ctx context.Context, client *gophercloud.ServiceClient) (r RuntimeInfoResult) {
	_ = "STUB: not implemented"
	return *new(RuntimeInfoResult)
}

// CleanTombstones removes deleted data from disk and cleans up the existing tombstones.
func CleanTombstones(ctx context.Context, client *gophercloud.ServiceClient) (r CleanTombstonesResult) {
	_ = "STUB: not implemented"
	return *new(CleanTombstonesResult)
}

// DeleteSeriesOptsBuilder allows extensions to add parameters to the DeleteSeries request.
type DeleteSeriesOptsBuilder interface {
	ToMetricDeleteSeriesQuery() (string, error)
}

// DeleteSeriesOpts contains the options for deleting time series.
type DeleteSeriesOpts struct {
	// Match is a list of series selectors. At least one must be provided.
	Match []string `q:"match[]" required:"true"`

	// Start is the start timestamp (RFC3339 or Unix timestamp).
	Start string `q:"start"`

	// End is the end timestamp (RFC3339 or Unix timestamp).
	End string `q:"end"`
}

// ToMetricDeleteSeriesQuery formats DeleteSeriesOpts into a query string.
func (opts DeleteSeriesOpts) ToMetricDeleteSeriesQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DeleteSeries deletes data for a selection of series in a time range.
func DeleteSeries(ctx context.Context, client *gophercloud.ServiceClient, opts DeleteSeriesOptsBuilder) (r DeleteSeriesResult) {
	_ = "STUB: not implemented"
	return *new(DeleteSeriesResult)
}

// Snapshot creates a snapshot of all current data into snapshots/<datetime>-<rand>
// under the TSDB's data directory and returns the directory as response.
func Snapshot(ctx context.Context, client *gophercloud.ServiceClient) (r SnapshotResult) {
	_ = "STUB: not implemented"
	return *new(SnapshotResult)
}
