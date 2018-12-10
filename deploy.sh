#!/bin/bash

P_USER=$1

if [[ -z "${P_USER}" ]]; then
  P_USER=user1
fi

for folder in bin/*;
do
echo $folder
lambda_name=`basename $folder | sed 's/{//' | sed 's/}//'`
echo $lambda_name
zip bin/$lambda_name.zip bin/$lambda_name

aws s3api put-object --bucket handsonbartender --key $P_USER/$lambda_name.zip --body bin/$lambda_name.zip --profile epf

done


## we omit https://docs.aws.amazon.com/cli/latest/reference/cloudformation/package.html and we do it by hand
## because it is more convenient for the workshop
echo "deploying...."

sam deploy --template-file ./sam.yml --stack-name $P_USER-bartender-sam-deploy --parameter-overrides User=$P_USER --capabilities CAPABILITY_IAM --profile epf

rm bin/*.zip