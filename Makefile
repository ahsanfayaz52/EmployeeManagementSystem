include Makefile.variables

.PHONY: format todo test check prepare
## prefix before other make targets to run in your local dev environment
local: | quiet
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet: # this is silly but shuts up 'Nothing to be done for `local`'
	@:

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id
	
format: prepare
	${DOCKRUN} bash ./scripts/format.sh

check: prepare format
	${DOCKRUN} bash ./scripts/check.sh

test:  check db_prepare
	${DOCKTEST} bash ./scripts/test.sh

db_start: db_stop
	@docker run --name employee-system-mysql-db -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306 -d mysql:5.6
	@docker run --name employee-employee-mongo-db -p 27015-27017:27015-27017 -d mongo:4.2.0

db_prepare: db_start
	@docker cp employee_system.sql employee-system-mysql-db:employee_system.sql
	@echo "Executing databases...wait for 15 seconds"
	@sleep 15
	@docker exec -i employee-system-mysql-db sh -c 'mysql -uroot < employee_system.sql'
	
db_stop:
	bash ./scripts/db_stop.sh
