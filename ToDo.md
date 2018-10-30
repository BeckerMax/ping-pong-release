src -> all source code from the job, can be git submodules packaged by package
jobs -> Describes to the agent how to start the release, templates and how to render them.
packages -> source code dependencies and package scripts


- bosh vendor-package golang-1.11-linux ~/workspace/github.com/bosh-packages/golang-release
- path to golang library the log dir so stdout is found in /var/sys/log/<job>
- Add a way to path in ip in ping job over manifest
  - do I need env variable or passing something to the binary?
- Provide the ip over bosh links from pong
- Write a release for the ruby license server
- try out writing a drain script
- Update create release documentation to include bpm
- Change Specify the property in the deployment manifest to `3.`


DEPLOY:
create bosh-lite
cd ~/deployments/virtualbox
../../workspace/bosh-deployment/virtualbox/create-env.sh

cd ~/max/ping-pong-release
. ../../deployments/virtualbox/.envrc
bosh upload-release
bosh ucc example-manifest/warden-cloud-config.yml

 bosh upload-stemcell --sha1 79833dcea376478a779cecd93b07394f2a363102 \
  https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-xenial-go_agent?v=97.28
