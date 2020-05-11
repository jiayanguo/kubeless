## Prerequisites:
1. install serverless `curl -o- -L https://slss.io/install | bash`

~~2. install serverless plugin `npm install -g serverless-kubeless`~~

## Steps
1. create a project

    `serverless create --template kubeless-python`
2. change runtime to go in serverless.yml 
3. change serverless-kubeless version. Old version will not work with go.

   ` "serverless-kubeless": "^0.7.0"`
3. create a `handler.go`
```
package kubeless

import (
	"github.com/kubeless/kubeless/pkg/functions"
)

// Foo sample function
func Handler(event functions.Event, context functions.Context) (string, error) {
	return "Hello world!", nil
}
```
4. call the function

`serverless invoke -f go-echo -l --data 'hello!'




