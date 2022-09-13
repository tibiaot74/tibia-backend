local-init:
	@docker-compose up -d
	@sleep 10
	echo "Ready to go!"
	@gin --port 3000 --appPort 7474

local-destroy:
	@docker-compose down