# DataCommons Mixer

Data Commons Mixer is an API server that serves Data Commons Data API. It gets deployed in a Kubernetes cluster.

## Git Development Process

In Mixer GitHub [repo](https://github.com/datacommonsorg/mixer), click on "Fork"
button to fork the repo.

Clone the forked repo to your local machine and add datacommonsorg/mixer repo as a remote.

```shell
git clone git@github.com:<YOUR-REPO>/mixer.git
cd mixer
git remote add dc https://github.com/datacommonsorg/mixer.git
```

Every time when you want to send a Pull Request, do the following steps:

```shell
git checkout master
git pull dc master
git checkout -b {new_local_branch_name}
# Make some code change
git add .
git commit -m "commit message"
git push -u origin {new_local_branch_name}
```

Then in your forked repo, can send a Pull Request. Wait for approval of the Pull Request and merge the change.

If you want to pull the latest remote code into your local branch that is under development, do:

```shell
git checkout {local_branch_name}
git pull dc master
git push -u origin {local_branch_name}
```

## Develop and test locally

Follow the [Developer Guide](docs/developer_guide.md).

## Setup a new GKE cluster

Follow GKE [Setup Guide](gke/README.md).
