---
title: Metrics
description: Learn what Keptn Metrics are and how to use them
icon: concepts
layout: quickstart
weight: 10
hidechildren: true # this flag hides all sub-pages in the sidebar-multicard.html
---

### Keptn Metric

A `KeptnMetric` is a CRD representing a metric.
The metric will be collected from the provider specified in the
specs.provider.name field.
The query is a string in the provider-specific query language, used to obtain a metric.
Providing the metrics as CRD into a K8s cluster will facilitate the reusability of this data across multiple components.
Furthermore, this allows using multiple observability platforms for different metrics.

A `KeptnMetric` looks like the following:

```yaml
apiVersion: metrics.keptn.sh/v1alpha1
kind: KeptnMetric
metadata:
  name: keptnmetric-sample
  namespace: podtato-kubectl
spec:
  provider:
    name: "prometheus"
  query: "sum(kube_pod_container_resource_limits{resource='cpu'})"
  fetchIntervalSeconds: 5
```

In this example, the provider is set to `prometheus`, which is one of the currently supported `KeptnMetricProviders`.
The provider tells the metrics-operator where to get the value for the `KeptnMetric`, and its configuration looks follows:

```yaml
apiVersion: metrics.keptn.sh/v1alpha3
kind: KeptnMetricsProvider
metadata:
  name: prometheus
spec:
  targetServer: "http://prometheus-k8s.monitoring.svc.cluster.local:9090"
```

Other supported providers are `dynatrace`, and `dql`:

````yaml
apiVersion: metrics.keptn.sh/v1alpha3
kind: KeptnMetricsProvider
metadata:
  name: dynatrace
  namespace: podtato-kubectl
spec:
  targetServer: "<dynatrace-tenant-url>"
  secretKeyRef:
    name: dt-api-token
    key: DT_TOKEN
---
apiVersion: metrics.keptn.sh/v1alpha3
kind: KeptnMetricsProvider
metadata:
  name: dql
  namespace: podtato-kubectl
spec:
  secretKeyRef:
    key: CLIENT_SECRET
    name: dt-third-gen-secret 
  targetServer: "<dynatrace-third-gen-target-server>"
````

Keptn metrics can be exposed as OTel metrics via port `9999` of the KLT metrics-operator.
To expose them, the env
variable `EXPOSE_KEPTN_METRICS` in the metrics-operator manifest needs to be set to `true`.
The default value of this variable
is `true`.
To access the metrics, use the following command:

```shell
kubectl port-forward deployment/metrics-operator 9999 -n keptn-lifecycle-toolkit-system
```

and access the metrics via your browser with:

```http://localhost:9999/metrics```

#### Accessing Metrics via the Kubernetes Custom Metrics API

`KeptnMetrics` can also be retrieved via the Kubernetes Custom Metrics API.
This makes it possible to refer to these metrics via the Kubernetes *HorizontalPodAutoscaler*, as in the following
example:

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: podtato-head-entry
  namespace: podtato-kubectl
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: podtato-head-entry
  minReplicas: 1
  maxReplicas: 10
  metrics:
    - type: Object
      object:
        metric:
          name: keptnmetric-sample
        describedObject:
          apiVersion: metrics.keptn.sh/v1alpha1
          kind: KeptnMetric
          name: keptnmetric-sample
        target:
          type: Value
          value: "10"
```

You can also use the `kubectl raw` command to retrieve the values of a `KeptnMetric`, as in the following example:

```shell
$ kubectl get --raw "/apis/custom.metrics.k8s.io/v1beta2/namespaces/podtato-kubectl/keptnmetrics.metrics.sh/keptnmetric-sample/keptnmetric-sample" | jq .

{
  "kind": "MetricValueList",
  "apiVersion": "custom.metrics.k8s.io/v1beta2",
  "metadata": {},
  "items": [
    {
      "describedObject": {
        "kind": "KeptnMetric",
        "namespace": "podtato-kubectl",
        "name": "keptnmetric-sample",
        "apiVersion": "metrics.keptn.sh/v1alpha1"
      },
      "metric": {
        "name": "keptnmetric-sample",
        "selector": {
          "matchLabels": {
            "app": "frontend"
          }
        }
      },
      "timestamp": "2023-01-25T09:26:15Z",
      "value": "10"
    }
  ]
}
```

You can also filter based on matching labels.
So to e.g. retrieve all metrics that are labelled with `app=frontend`, you
can use the following command:

```shell
$ kubectl get --raw "/apis/custom.metrics.k8s.io/v1beta2/namespaces/podtato-kubectl/keptnmetrics.metrics.sh/*/*?labelSelector=app%3Dfrontend" | jq .

{
  "kind": "MetricValueList",
  "apiVersion": "custom.metrics.k8s.io/v1beta2",
  "metadata": {},
  "items": [
    {
      "describedObject": {
        "kind": "KeptnMetric",
        "namespace": "keptn-lifecycle-toolkit-system",
        "name": "keptnmetric-sample",
        "apiVersion": "metrics.keptn.sh/v1alpha1"
      },
      "metric": {
        "name": "keptnmetric-sample",
        "selector": {
          "matchLabels": {
            "app": "frontend"
          }
        }
      },
      "timestamp": "2023-01-25T09:26:15Z",
      "value": "10"
    }
  ]
}
```
