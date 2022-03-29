.PHONY: up restart logs build stage prod

repo=vivaconagua/pool-core



up:
	docker-compose up -d

restart:
	docker-compose restart

logs:
	docker-compose logs app

build:
	docker-compose -f docker-compose.build.yml build --force-rm --no-cache

stage:
	docker push ${repo}:stage

prod:
	docker tag ${repo}:stage ${repo}:latest
	docker push ${repo}:latest
