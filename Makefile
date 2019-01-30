.PHONY: build clean deploy
BOT_KEY: ${BOT_KEY}
BOT_KEY_SECRET: ${BOT_KEY_SECRET}
BOT_TOKEN: ${BOT_TOKEN}
BOT_TOKEN_SECRET: ${BOT_TOKEN_SECRET}

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/tweet tweet/tweet.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
