VERSION := "1.0.0"
FLAGS 	:= "-s -w -X=main.build=`git rev-parse --short HEAD` -X=main.version=$(VERSION)"
REPO 		:= sockwait
USER 		:= yawn
TOKEN 	= `cat .token`

sw:
	gox build -ldflags $(FLAGS) sw.go

build:
	mkdir -p build
	gox -osarch="linux/amd64 linux/arm darwin/amd64" -ldflags $(FLAGS) -output "build/{{.OS}}-{{.Arch}}/sw"
	find build -name sw -exec upx --brute -q {} +

clean:
	rm -rf build/

release:
	git tag $(VERSION) -f && git push --tags -f
	github-release release --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN)
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name sw-osx --file build/darwin/sw
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name sw-linux --file build/linux-amd64/sw
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name sw-linux-arm --file build/linux-arm/sw
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name sw-osx --file build/darwin-amd64/sw

retract:
	github-release delete --tag $(VERSION) -s $(TOKEN)

.PHONY: build clean release retract
