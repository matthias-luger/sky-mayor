# generates the c# client
VERSION=0.3.2
NAME=Coflnet.Sky.Mayor.Client

docker run --rm -v "${PWD}:/local" --network host -u $(id -u ${USER}):$(id -g ${USER})  openapitools/openapi-generator-cli generate \
-i /local/swagger.json \
-g csharp \
-o /local/out --additional-properties=packageName=$NAME,packageVersion=$VERSION,licenseId=MIT,targetFramework=net6.0

cd out
path=src/$NAME/$NAME.csproj
sed -i 's/GIT_USER_ID/Coflnet/g' $path
sed -i 's/GIT_REPO_ID/sky-mayor/g' $path
sed -i 's/>OpenAPI/>Coflnet/g' $path
sed -i 's@annotations</Nullable>@annotations</Nullable>\n    <PackageReadmeFile>README.md</PackageReadmeFile>@g' $path
sed -i 's@Remove="System.Web" />@Remove="System.Web" />\n    <None Include="../../../../README.md" Pack="true" PackagePath="\"/>@g' $path

dotnet pack
cp src/$NAME/bin/Release/$NAME.*.nupkg ..
