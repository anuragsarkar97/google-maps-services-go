package metrics_test

import (
	"context"
	maps "github.com/anuragsarkar97/google-maps-services-go"
	"github.com/anuragsarkar97/google-maps-services-go/metrics"
	"go.opencensus.io/stats/view"
	"testing"
)

func TestClientWithOpenCensus(t *testing.T) {
	metrics.RegisterViews()
	server := mockServer([]int{200, 400}, `{"results" : [], "status" : "OK"}`)
	defer server.Close()
	c, err := maps.NewClient(
		maps.WithAPIKey("AIza-Maps-API-Key"),
		maps.WithBaseURL(server.URL),
		maps.WithMetricReporter(metrics.OpenCensusReporter{}))
	if err != nil {
		t.Errorf("Unable to create client with OpenCensusReporter")
	}
	r := &maps.ElevationRequest{
		Locations: []maps.LatLng{
			{
				Lat: 39.73915360,
				Lng: -104.9847034,
			},
		},
	}
	_, err = c.Elevation(context.Background(), r)
	if err != nil {
		t.Errorf("r.Get returned non nil error, was %+v", err)
	}
	_, err = c.Elevation(context.Background(), r)
	if err != nil {
		t.Errorf("r.Get returned non nil error, was %+v", err)
	}
	count, _ := view.RetrieveData("maps.googleapis.com/client/count")
	if len(count) != 2 {
		t.Errorf("expected two metrics, got %v", len(count))
	}
}
