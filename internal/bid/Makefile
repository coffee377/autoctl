TIMESTAMP := $(shell date +%Y%m%d%H%M%S)
DEPT ?= 1

dump-data:
	@cat data/bid_project.sql | sed 's/cds_infra\.//g' > migrations/20251113082500_data.sql; \
	echo "\n" >> migrations/20251113082500_data.sql; \
	cat data/bid_apply.sql | sed 's/cds_infra\.//g' >> migrations/20251113082500_data.sql; \
	echo "\n" >> migrations/20251113082500_data.sql; \
	cat data/bid_expense.sql | sed 's/cds_infra\.//g' >> migrations/20251113082500_data.sql; \
	echo "\n" >> migrations/20251113082500_data.sql; \
	cat data/bid_info.sql | sed 's/cds_infra\.//g' >> migrations/20251113082500_data.sql; \
	echo "\n" >> migrations/20251113082500_data.sql
	@make hash

dump-opts:
	@mysqldump -h 127.0.0.1 -P 3306 -u root --password='root!@@&' \
		cds_infra \
		bid_apply_opts bid_expense_opts bid_info_opts \
		> bid_opts.sql

clean:
	@atlas schema clean --env dev --var db_name=ccl_base

diff: hash
	@export GOWORK=off; \
	 atlas migrate diff --env dev

apply:
	@atlas migrate apply --env dev

apply-local:
	@atlas migrate apply 1 --env dev --var db_name=ccl_base

apply-test:
	@atlas migrate apply 1 --env "test" --allow-dirty \
	    --var db_host=10.1.83.26 \
	    --var db_port='3306' \
	    --var db_user=teamwork \
	    --var db_password='jqkj5350**)' \
	    --var db_name=teamwork

status:
	atlas migrate status --env dev

status-local:
	atlas migrate status --env dev --var db_name=ccl_base

status-test:
	@atlas migrate status --env 'test' \
	    --var db_host=10.1.83.26 \
	    --var db_port='3306' \
	    --var db_user=teamwork \
	    --var db_password='jqkj5350**)' \
	    --var db_name=teamwork

hash:
	@atlas migrate hash --dir file://migrations

create: hash
	@atlas migrate new --dir file://migrations