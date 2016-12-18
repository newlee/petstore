How to use:
Run once:

$ vagrant up

Then run script when change code:

$ vagrant ssh -c 'cd /vagrant && ./build/build.sh'

And open http://localhost:8090/ in a web browser.

Save currently-used dependencies to file Godeps:

$ godep save ./...

Build project using saved dependencies:

$ godep go install ./...