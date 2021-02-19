.DEFAULT_GOAL := generate

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

generate: generate_go_docker

generate_go_docker:
	docker run --rm -v $(current_dir):/defs namely/protoc-all -d protos/** -l go

push_all:
	docker build . -f docker/auth.dockerfile -t gcr.io/keeping-it-casual/auth
	docker push gcr.io/keeping-it-casual/auth

	docker build . -f docker/server.dockerfile -t gcr.io/keeping-it-casual/server
	docker push gcr.io/keeping-it-casual/server