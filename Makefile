.PHONY: up restart logs build stage prod

repo=vivaconagua/pool-event

pre-commit:
	pre-commit run --show-diff-on-failure --color=always --all-files

commit:
	pre-commit run --show-diff-on-failure --color=always --all-files && git commit && git push

test:
	go test ./dao && go test ./handlers/token && go test ./handlers/admin

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
