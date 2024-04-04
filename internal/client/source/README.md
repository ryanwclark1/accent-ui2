docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-auth.yml \
  -g go \
  --additional-properties=generateInterfaces=true,packageName=accentauth,packageVersion=0.0.1,prependFormOrBodyParameters=true,enumClassPrefix=true,isGoSubmodule=true,structPrefix=true \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/client/accentauth

docker run --rm \
-v ${PWD}:/local openapitools/openapi-generator-cli validate \
-i /local/out/client/accentauth/api/openapi.yaml

docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-auth.yml \
  -g go-server \
  --additional-properties=packageName=accentauthserver,packageVersion=0.0.1,router=mux,sourceFolder=src \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/server/accentauth

addResponseHeaders=true
emumClassPrefix=true
featureCORS=false
hideGenerationTimestamp=true
onlyInterfaces=false
outputAsLibrary=true
router=mux
sourceFolder=src


docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-auth.yml \
  -g graphql-schema \
  --additional-properties=packageName=accentauthgraphql,packageVersion=0.0.1 \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/graphql/accentauth


docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-auth.yml \
  -g mysql-schema \
  --additional-properties=defaultDatabaseName=accent,identifierNamingConvention=original,jsonDataTypeEnabled=true,namedParametersEnabled=false \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/mysql/accentauth

docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-auth.yml \
  -g postman-collection \
  --additional-properties=folderStrategy=true,pathParamsAsVariables=true,postmanGuid=true,poastmanGuidPlaceholderName=true \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/postman/accentauth


docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-auth.yml \
  -g protobuf-schema \
  --additional-properties=numberedFieldNumberList=false,startEnumsWithUnknown=false \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/protobuf/accentauth

docker run \
-v ${PWD}:/workdir davidanson/markdownlint-cli2 --fix "/workdir/out/client/accentauth/docs/*.md" "#node_modules"


docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/accent-call-logd.yml \
  -g go \
  --additional-properties=packageName=accentcalllogd,packageVersion=0.0.1,prependFormOrBodyParameters=true,enumClassPrefix=true,isGoSubmodule=true,structPrefix=true,generateInterfaces=true \
  --git-user-id ryanwclark \
  --git-repo-id accent-voice \
  -o /local/out/client/accentcalllogd


var defaultValue []int32 = [1,2,3,4,5,6,7]
var defaultValue []int32 = []int32{1,2,3,4,5,6,7}
