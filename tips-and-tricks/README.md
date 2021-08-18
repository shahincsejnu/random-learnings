# Tips & Tricks

## How to span workspace in multiple monitors in ubuntu 20.04?
- Install `gnome-tweaks` by :
  - `sudo apt install gnome-tweaks`
- Then use `gnome-tweaks` commands (will prompt up the tweaks box):
  - go to the options, specially the last one (`Workspaces`) and tick the `Workspaces span displays` option
- Now, your workspace will span in all of your monitors


## How to use private repo in go.mod? 
- To use private repo in go.mod and then to do `go mod tidy; go mod vendor` follow the below instructions:
  - set `GOPRIVATE` env variable : `go env -w GOPRIVATE=<repo_name>`
    - Ex: `go env -w GOPRIVATE=kubeform.dev/<repo_name>`
    - Note that do not use `export` to set the variable, use `go env -w`:
      - cause doing exporting will not work fine and if you do both export and go env then it will get conflict (faced that issue)
    - also add this one : `git config --global --add url.git@github.com:.insteadOf https://github.com/` :
      - To tell to use ssh rather than https