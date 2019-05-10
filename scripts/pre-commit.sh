#!/bin/sh

GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

if [[ "$GO_FILES" = "" ]]; then
  exit 0
fi

GOLINT=$GOPATH/bin/golangci-lint
GOIMPORTS=$GOPATH/bin/goimports
APP_DIR=app/*.go

# Check for golint
if [[ ! -x "$GOLINT" ]]; then
  printf "\t\033[41mPlease install golint in your local machine thru\033[0m (go get -u github.com/golangci/golangci-lint/cmd/golangci-lint)"
  exit 1
fi

# Check for goimports
if [[ ! -x "$GOIMPORTS" ]]; then
  printf "\t\033[41mPlease install goimports\033[0m (go get golang.org/x/tools/cmd/goimports)"
  exit 1
fi

PASS=true

for FILE in $GO_FILES
do
  # Run goimports on the staged file
  $GOIMPORTS -w $FILE

  # Run golint on the staged file and check the exit status
	echo "LINTING ->" $FILE
  $GOLINT run $APP_DIR
  if [[ $? == 1 ]]; then
    printf "\t\033[31mgolint $FILE\033[0m \033[0;30m\033[41mFAILURE!\033[0m\n"
    PASS=false
  else
    printf "\t\033[32mgolint $FILE\033[0m \033[0;30m\033[42mpass\033[0m\n"
  fi

  # Run govet on the staged file and check the exit status
	echo "RUNNING GO VET -> " $FILE
  go vet $APP_DIR
  if [[ $? != 0 ]]; then
    printf "\t\033[31mgo vet $FILE\033[0m \033[0;30m\033[41mFAILURE!\033[0m\n"
    PASS=false
  else
    printf "\t\033[32mgo vet $FILE\033[0m \033[0;30m\033[42mpass\033[0m\n"
  fi
done

if ! $PASS; then
  printf "\033[0;30m\033[41mCOMMIT FAILED PLEASE CHECK YOUR WORK\033[0m\n"
  exit 1
else
  printf "\033[0;30m\033[42mCOMMIT SUCCEEDED\033[0m\n"
fi

exit 0