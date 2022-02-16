let responseBody="";
const jsonData = JSON.parse(responseBody);

var infos = jsonData.body.infos

for (var k in infos) {
    console.log(infos[k])

}