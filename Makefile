
PROJECT_ID = 0daryo/cinemawith

init:
	@echo Initialize cinemawith now......
	go get -u github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	@echo Initialize cinemawith completed!!!!

build:
	docker-compose build

start:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs api

test:
	go test -test.v ./src/...

mockgen_task:
	$(eval SERVICE_LIST := $(call get_service_list))	
	$(foreach file, $(SERVICE_LIST), $(call mockgen_service,$(shell basename $(file))))
	$(eval REPOSITORY_LIST := $(call get_repository_list))
	$(foreach file, $(REPOSITORY_LIST), $(call mockgen_repository,$(shell basename $(file))))

define get_service_list
	$(shell	find ./src/service/ -maxdepth 1 -type f ! -name "*_impl*.go")
endef

define mockgen_service
	$(shell mockgen -source ./src/service/$1 -destination ./src/service/mock/$1)
endef

define get_repository_list
	$(shell	find ./src/domain/repository -maxdepth 1 -type f )
endef

define mockgen_repository
	$(shell mockgen -source ./src/domain/repository/$1 -destination ./src/domain/repository/mock/$1)
endef