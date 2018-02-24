BUILD=main
DEPLOYMENT=$(BUILD).zip
FUNCTION=Passport
REGION=ap-southeast-1
FunctionName=Passport
BIN = main

all: deploy
	echo $(FunctionName) 
	make git

build: ($BUILD)
	echo $(FunctionName)
#	go build main.go
	GOOS=linux go build -o $(BUILD) .

package:
	zip main.zip main

deploy: package
    aws lambda update-function-code \
        --region $(REGION) \
        --function-name $(FUNCTION) \
        --zip-file fileb://./$(DEPLOYMENT)

test:
	aws lambda invoke --function-name  $(FunctionName) lambda.log

git:
	make clean
	git add . 
	git commit -m "auot make update"
	git push

.PHONY: clean
clean:
	rm -f $(BUILD) $(DEPLOYMENT)
	rm -f lambda.log
