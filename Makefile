local-init:
	@docker-compose -f docker-compose.yaml up -d
	@sleep 10
	echo "Ready to go!"
	@gin --port 3000 --appPort 7474

local-destroy:
	@docker-compose -f docker-compose.yaml down

test:
	echo "Setting up the test environment..."
	@docker-compose -f docker-compose-test.yaml up --build

test-destroy:
	@docker-compose -f docker-compose-test.yaml down