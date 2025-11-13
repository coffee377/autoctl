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
	@atlas schema clean --env dev

diff: hash
	@export GOWORK=off; \
	 atlas migrate diff --env dev

apply:
	@atlas migrate apply --env dev

status:
	atlas migrate status --env dev

hash:
	@atlas migrate hash --dir file://migrations

create: hash
	@atlas migrate new --dir file://migrations