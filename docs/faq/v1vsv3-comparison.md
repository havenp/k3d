# Feature Comparison: v1 vs. v3

## v1.x feature -> implementation in v3

```text
- k3d
  - check-tools                   -> won't do
  - shell                         -> planned: `k3d shell CLUSTER`
    - --name                      -> planned: drop (now as arg)
    - --command                   -> planned: keep
    - --shell                     -> planned: keep (or second arg)
      - auto, bash, zsh
  - create                        -> `k3d create cluster CLUSTERNAME`
    - --name                      -> dropped, implemented via arg
    - --volume                    -> implemented
    - --port                      -> implemented
    - --port-auto-offset          -> TBD
    - --api-port                  -> implemented
    - --wait                      -> implemented
    - --image                     -> implemented
    - --server-arg                -> implemented as `--k3s-server-arg`
    - --agent-arg                 -> implemented as `--k3s-agent-arg`
    - --env                       -> planned
    - --label                     -> planned
    - --workers                   -> implemented
    - --auto-restart              -> dropped (docker's `unless-stopped` is set by default)
    - --enable-registry           -> planned (possible consolidation into less registry-related commands?)
    - --registry-name             -> TBD
    - --registry-port             -> TBD
    - --registry-volume           -> TBD
    - --registries-file           -> TBD
    - --enable-registry-cache     -> TBD
  - (add-node)                    -> `k3d create node NODENAME`
    - --role                      -> implemented
    - --name                      -> dropped, implemented as arg
    - --count                     -> implemented as `--replicas`
    - --image                     -> implemented
    - --arg                       -> planned
    - --env                       -> planned
    - --volume                    -> planned
    - --k3s                       -> TBD
    - --k3s-secret                -> TBD
    - --k3s-token                 -> TBD
  - delete                        -> `k3d delete cluster CLUSTERNAME`
    - --name                      -> dropped, implemented as arg
    - --all                       -> implemented
    - --prune                     -> TBD
    - --keep-registry-volume      -> TBD
  - stop                          -> `k3d stop cluster CLUSTERNAME`
    - --name                      -> dropped, implemented as arg
    - --all                       -> implemented
  - start                         -> `k3d start cluster CLUSTERNAME`
    - --name                      -> dropped, implemented as arg
    - --all                       -> implemented
  - list                          -> dropped, implemented as `k3d get clusters`
  - get-kubeconfig                -> `k3d get kubeconfig CLUSTERNAME`
    - --name                      -> dropped, implemented as arg
    - --all                       -> implemented
    - --overwrite                 -> implemented
  - import-images                 -> `k3d load image [--cluster CLUSTERNAME] [--keep] IMAGES`
    - --name                      -> implemented as `--cluster`
    - --no-remove                 -> implemented as `--keep`
```
