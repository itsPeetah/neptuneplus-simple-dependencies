IMAGE_BASE_NAME := itspeetah/neptuneplus-simple-dependencies
FUNCTIONS := waiter sequential parallel
WAITERS := 1 2 3 4 5
CALLERS := sequential parallel

PYTHON := $(shell which python3)
MANIFEST_SCRIPT := ./codegen/yamlgen.py

buildall: $(addprefix build-,$(FUNCTIONS))

build-%:
	@echo "Building Docker image for function $*..."
	docker build --target neptuneplus-simple-dependencies-$* -t $(IMAGE_BASE_NAME)-$*:latest .
	@echo "Successfully built $(IMAGE_BASE_NAME)-$*:latest"

clean:
	@echo "Removing Docker images..."
	@for func in $(FUNCTIONS); do \
		if docker images -q $(IMAGE_BASE_NAME)-$$func:latest &> /dev/null; then \
			docker rmi $(IMAGE_BASE_NAME)-$$func:latest; \
			echo "Removed $(IMAGE_BASE_NAME)-$$func:latest"; \
		else \
			echo "Image $(IMAGE_BASE_NAME)-$$func:latest does not exist."; \
		fi \
	done
	@echo "Clean up complete."

pushall: $(addprefix push-,$(FUNCTIONS))

push-%:
	@echo "Deploying Docker imgage for function $*..."
	docker image push $(IMAGE_BASE_NAME)-$*:latest
	@echo "Successfully pushed $(IMAGE_BASE_NAME)-$*:latest"

manifests: $(addprefix manifest-function-waiter-,$(WAITERS)) manifest-function-caller-sequential manifest-function-caller-parallel

manifest-function-waiter-%:
	@echo "Generating manifest for function waiter-$*"
	$(PYTHON) $(MANIFEST_SCRIPT) waiter-$* $(IMAGE_BASE_NAME)-waiter
	@echo "Done"

manifest-function-caller-%:
	@echo "Generating manifest for function caller-$*"
		$(PYTHON) $(MANIFEST_SCRIPT) caller-$* $(IMAGE_BASE_NAME)-$*
	@echo "Done"

release: buildall pushall

deploy:
	kubectl apply -f ./config/deploy

undeploy:
	kubectl delete -f ./config/deploy

redeploy: undeploy deploy