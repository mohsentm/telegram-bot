#Check which operating system we are having and use specified command to build it
OSFLAG :=
BUILDSTART :=

build:
	$(BUILDSTART) go build -o ./bin/telebot ./cmd/
clean:
	$(CLEAN)

.PHONY: build clean