# k8synth

> :construction: work in progress

Your k8s cluster is your favorite noise artist.

## Install

Only using `go` (>=1.21) for now:
```
go install
```

## Usage

`k8synth` triggers midi signals. You can use a midi synth to hear something. For example:
[https://neauoire.github.io/Enfer/](https://neauoire.github.io/Enfer/)

```
// Select you k8s namespace
k8synth -namepsace=mynamespace

// Optionam. Absolute path to your kube config. Default to the usual.
k8synth -kubeconfig=~/.kube/config
```