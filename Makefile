PID_FILE = /tmp/car-sales-api.pid

prepare:
	go mod download

build:
	go build -o bin/api main.go

kill:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 

run:
	clear && go run main.go & echo $$! > $(PID_FILE)

watch: run
	fswatch -or --event=Updated -e ".git" . | \
	xargs -n1 make reload

reload: kill watch

# .PHONY is used for reserving tasks words
.PHONY: build kill reload watch prepare run

# supress echo commands on cli
.SILENT: build kill run watch prepare