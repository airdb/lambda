BUILD=main
DEPLOYMENT=$(BUILD).zip
FUNCTION=Passport
REGION=ap-northeast-1

all:
	echo $(Function) auto make success!!!
#	make git
	aws lambda update-function-code \
        --region $(REGION) \
        --function-name $(FUNCTION) \
        --zip-file fileb://./$(DEPLOYMENT)

build: 
#	echo $(FunctionName)
#	go build main.go
	GOOS=linux go build -o $(BUILD) .

package: build
	zip main.zip main

test:
	aws lambda invoke --function-name $(FunctionName) lambda.log

git:
	make clean
	git add . 
	git commit -m "auto make update"
	git push

.PHONY: clean
clean:
	rm -f $(BUILD) $(DEPLOYMENT)
	rm -f lambda.log
