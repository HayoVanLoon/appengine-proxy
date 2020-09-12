
check:
ifndef PROJECT
	$(error missing PROJECT)
endif

deploy: check
	cp app.yaml app.yaml.example
	cp etc/app.yaml .
	gcloud app deploy \
		--project=$(PROJECT)
	cp app.yaml.example app.yaml
