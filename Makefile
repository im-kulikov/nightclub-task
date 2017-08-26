.PHONY: main

main: clean
	set -x \
		&& GOOS=darwin GOARCH=amd64 gb build --tags macOS \
		&& mv ./bin/club-darwin-amd64-macOS ./bin/club

clean:
	rm -rf ./bin/*

race: clean
	set -x \
		&& GOOS=darwin GOARCH=amd64 gb build -race --tags macOS \
		&& mv ./bin/club-darwin-amd64-macOS-race ./bin/club \
		&& ./bin/club
