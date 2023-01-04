IMAGE_NAME := dbt_with_ch

.DEFAULT_GOAL := dbt_version
.PHONY: dbt_version

dbt:
	docker build --no-cache -t $(IMAGE_NAME) -f Docker/dbt.Dockerfile .


dbt_version:
	docker run --rm -it $(IMAGE_NAME) dbt --version

dbt_run:
	cd dbt && ../bin/edbt run --select example --profile ch_docker
