FunctionName=Passport
BIN = main

all: 
	echo $(FunctionName)
	go build main.go
	zip main.zip main
	aws lambda update-function-code --function-name $(FunctionName) --zip-file fileb://main.zip
	aws lambda invoke --function-name  $(FunctionName) lambda.log

git:
	make clean
	git add . 
	git commit -m "auot make update"
	git push


clean:
	rm -f lambda.log
	rm -f main.zip
	rm -f $(BIN)

