build:
	docker build . -t mao-ctrl

run:
	docker run -it -d --rm -p 3000:3000 --env-file configs/.env mao-ctrl
