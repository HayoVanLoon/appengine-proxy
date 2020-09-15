
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
