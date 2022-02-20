.PHONY: all
all: mugambo

.PHONY: mugambo
mugambo:
	GIT_COMMIT=`git rev-list -1 HEAD 2>/dev/null || echo ""` && \
	GIT_DATE=`git log -1 --date=short --pretty=format:%ct 2>/dev/null || echo ""` && \
	go build \
	    -ldflags "-s -w -X github.com/topcoder1208/fantom-fork/cmd/mugambo/launcher.gitCommit=$${GIT_COMMIT} -X github.com/topcoder1208/fantom-fork/cmd/mugambo/launcher.gitDate=$${GIT_DATE}" \
	    -o build/mugambo \
	    ./cmd/mugambo

TAG ?= "latest"
.PHONY: mugambo-image
mugambo-image:
	docker build \
    	    --network=host \
    	    --build-arg GOPROXY=$(GOPROXY) \
    	    -f ./docker/Dockerfile.mugambo -t "mugambo:$(TAG)" .

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -coverprofile=cover.prof $$(go list ./... | grep -v '/gossip/contract/' | grep -v '/gossip/emitter/mock' | xargs)
	go tool cover -func cover.prof | grep -e "^total:"

.PHONY: fuzz
fuzz:
	CGO_ENABLED=1 \
	mkdir -p ./fuzzing && \
	go run github.com/dvyukov/go-fuzz/go-fuzz-build -o=./fuzzing/gossip-fuzz.zip ./gossip && \
	go run github.com/dvyukov/go-fuzz/go-fuzz -workdir=./fuzzing -bin=./fuzzing/gossip-fuzz.zip


.PHONY: clean
clean:
	rm -fr ./build/*
