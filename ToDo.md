src -> all source code from the job, can be git submodules packaged by package
scripts
packages -> source code dependencies and package scripts


- build binary
- packaging
- path to golang library the log dir so stdout is found in /var/sys/log/<job>
- Add a way to path in ip in ping job over manifest
  - do I need env variable or passing something to the binary?
- Provide the ip over bosh links from pong
- Write a release for the ruby license server
- try out writing a drain script
- Update create release to include bpm
