TEST_ARGS ?= --count=1

.PHONY: init
init:
	@rm go.work go.work.sum
	@go work init
	@go work use -r .
	@go work sync

#.PHONY: build_libgit2
#build_libgit2:
#	@./script/build_libgit2.sh

# Bundled dynamic library
# =======================
# In order to avoid having to manipulate `git_dynamic.go`, which would prevent
# the system-wide libgit2.so from being used in a sort of ergonomic way, this
# instead moves the complexity of overriding the paths so that the built
# libraries can be found by the build and tests.
.PHONY: build-libgit2-dynamic
build-libgit2-dynamic:
	@./script/build-libgit2-dynamic.sh

# Bundled static library
# ======================
# This is mostly used in tests, but can also be used to provide a
# statically-linked library with the bundled version of libgit2.
.PHONY: build-libgit2-static
build-libgit2-static:
	@./script/build-libgit2-static.sh

static-build/install/lib/libgit2.a:
	./script/build-libgit2-static.sh

test-static: static-build/install/lib/libgit2.a
	#go run script/check-MakeGitError-thread-lock.go
	go test --tags "static" $(TEST_ARGS) ./...

install-static: static-build/install/lib/libgit2.a
	go install --tags "static" ./...

#init:
#	git submodule add -f https://github.com/libgit2/libgit2.git vendor/libgit2

webhook:
	mc admin config set cds-test notify_webhook:file_download endpoint="http://10.1.42.244:9098/minio-events"
	mc admin service restart cds-test

event:
	mc event add cds-test/download arn:minio:sqs::file_download:webhook --event put --suffix .xlsx
	#mc event rm local/download arn:minio:sqs::primary:webhook --event put,delete --suffix .xlsx