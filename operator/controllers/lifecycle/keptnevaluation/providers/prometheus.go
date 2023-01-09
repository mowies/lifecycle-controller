package providers

import (
	"context"
	"fmt"      //nolint:gci
	"net/http" //nolint:gci
	"time"

	"github.com/go-logr/logr"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	promapi "github.com/prometheus/client_golang/api"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type KeptnPrometheusProvider struct {
	Log        logr.Logger
	httpClient http.Client
}

func (r *KeptnPrometheusProvider) EvaluateQuery(ctx context.Context, objective klcv1alpha2.Objective, provider klcv1alpha2.KeptnEvaluationProvider) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	queryTime := time.Now().UTC()
	r.Log.Info("Running query: /api/v1/query?query=" + objective.Query + "&time=" + queryTime.String())
	client, err := promapi.NewClient(promapi.Config{Address: provider.Spec.TargetServer, Client: &r.httpClient})
	if err != nil {
		return "", err
	}
	api := prometheus.NewAPI(client)
	result, w, err := api.Query(
		ctx,
		objective.Query,
		queryTime,
		[]prometheus.Option{}...,
	)

	if err != nil {
		return "", err
	}

	if len(w) != 0 {
		r.Log.Info("Prometheus API returned warnings: " + w[0])
	}

	// check if we can cast the result to a vector, it might be another data struct which we can't process
	resultVector, ok := result.(model.Vector)
	if !ok {
		return "", fmt.Errorf("could not cast result")
	}

	// We are only allowed to return one value, if not the query may be malformed
	// we are using two different errors to give the user more information about the result
	if len(resultVector) == 0 {
		r.Log.Info("No values in query result")
		return "", fmt.Errorf("no values in query result")
	} else if len(resultVector) > 1 {
		r.Log.Info("Too many values in the query result")
		return "", fmt.Errorf("too many values in the query result")
	}
	value := resultVector[0].Value.String()
	return value, nil
}
