# `kubectl krew neat`

- To see the output yaml more briefly, ex:
  - without : `kubectl get pods <pod_name> -o yaml`
    - output will contain full details including some unnecessary stuffs
  - with: `kubectl get pod <pod_name> -o yaml | kubectl neat`
    - output is now more precise and brief

- To use this we need to first install `kubectl krew`, to do that follow :
  - [install kubectl krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)
  - after installing `kubectl krew` we need to install `neat` plugin, by :
    - `kubectl krew install neat`