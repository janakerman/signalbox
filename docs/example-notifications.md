```
{
        "involvedObject": {
            "kind": "Kustomization",
            "namespace": "flux-system",
            "name": "podinfo",
            "uid": "624ec8a5-45c1-4450-9908-f6c2a1fa5eea",
            "apiVersion": "kustomize.toolkit.fluxcd.io/v1beta2",
            "resourceVersion": "28586"
        },
        "severity": "info",
        "timestamp": "2021-12-19T16:16:24Z",
        "message": "Reconciliation finished in 94.0083ms, next run in 30s",
        "reason": "ReconciliationSucceeded",
        "metadata": {
            "commit_status": "update",
            "revision": "master/132f4e719209eb10b9485302f8593fc0e680f4fc"
        },
        "reportingController": "kustomize-controller",
        "reportingInstance": "kustomize-controller-5db6bfc56d-tvx28"
    }
}
```
```
{
        "involvedObject": {
            "kind": "GitRepository",
            "namespace": "flux-system",
            "name": "podinfo",
            "uid": "708ae727-68cb-48ae-922f-fb06162be279",
            "apiVersion": "source.toolkit.fluxcd.io/v1beta1",
            "resourceVersion": "28730"
        },
        "severity": "info",
        "timestamp": "2021-12-19T16:16:36Z",
        "message": "Fetched revision: v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9",
        "reason": "info",
        "reportingController": "source-controller",
        "reportingInstance": "source-controller-565f8fbbff-j4sjz"
    }
}
```
```
{
        "involvedObject": {
            "kind": "Kustomization",
            "namespace": "flux-system",
            "name": "podinfo",
            "uid": "624ec8a5-45c1-4450-9908-f6c2a1fa5eea",
            "apiVersion": "kustomize.toolkit.fluxcd.io/v1beta2",
            "resourceVersion": "28691"
        },
        "severity": "error",
        "timestamp": "2021-12-19T16:16:36Z",
        "message": "kustomization path not found: stat /tmp/podinfo637577731/kustomize: no such file or directory",
        "reason": "ArtifactFailed",
        "metadata": {
            "revision": "v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9"
        },
        "reportingController": "kustomize-controller",
        "reportingInstance": "kustomize-controller-5db6bfc56d-tvx28"
    }
}
```
```
{
        "involvedObject": {
            "kind": "GitRepository",
            "namespace": "flux-system",
            "name": "podinfo",
            "uid": "708ae727-68cb-48ae-922f-fb06162be279",
            "apiVersion": "source.toolkit.fluxcd.io/v1beta1",
            "resourceVersion": "30398"
        },
        "severity": "info",
        "timestamp": "2021-12-19T16:24:24Z",
        "message": "Fetched revision: 6.0.3/ea292aa958c5e348266518af2261dc04d6270439",
        "reason": "info",
        "reportingController": "source-controller",
        "reportingInstance": "source-controller-565f8fbbff-j4sjz"
    }
}
```
```
{
        "involvedObject": {
            "kind": "Kustomization",
            "namespace": "flux-system",
            "name": "podinfo",
            "uid": "624ec8a5-45c1-4450-9908-f6c2a1fa5eea",
            "apiVersion": "kustomize.toolkit.fluxcd.io/v1beta2",
            "resourceVersion": "30303"
        },
        "severity": "info",
        "timestamp": "2021-12-19T16:24:25Z",
        "message": "Reconciliation finished in 132.9038ms, next run in 30s",
        "reason": "ReconciliationSucceeded",
        "metadata": {
            "commit_status": "update",
            "revision": "6.0.3/ea292aa958c5e348266518af2261dc04d6270439"
        },
        "reportingController": "kustomize-controller",
        "reportingInstance": "kustomize-controller-5db6bfc56d-tvx28"
    }
}
```
```
{
        "involvedObject": {
            "kind": "Kustomization",
            "namespace": "flux-system",
            "name": "podinfo",
            "uid": "624ec8a5-45c1-4450-9908-f6c2a1fa5eea",
            "apiVersion": "kustomize.toolkit.fluxcd.io/v1beta2",
            "resourceVersion": "30410"
        },
        "severity": "info",
        "timestamp": "2021-12-19T16:24:29Z",
        "message": "Reconciliation finished in 88.8099ms, next run in 30s",
        "reason": "ReconciliationSucceeded",
        "metadata": {
            "commit_status": "update",
            "revision": "6.0.3/ea292aa958c5e348266518af2261dc04d6270439"
        },
        "reportingController": "kustomize-controller",
        "reportingInstance": "kustomize-controller-5db6bfc56d-tvx28"
    }
}
```