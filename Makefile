
FILES := go.mod go.sum server.go

check:
ifndef PROJECT
	$(error missing PROJECT)
endif

deploy: check
	cp app.yaml app.yaml.old
	cp etc/app-$(PROJECT).yaml ./app.yaml
	-gcloud app deploy \
		--project=$(PROJECT)
	mv app.yaml.old app.yaml

zip:
	-rm -rf build
	mkdir -p build
	cp $(FILES) build
	cd build && go mod vendor
	-rm -rf out
	mkdir -p out
	cd build && zip -r ../out/appengine-cors-proxy.zip *
