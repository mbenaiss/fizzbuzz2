The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus question :
- Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request


## 🛠 Configuration
To update the configuration, you can do this by setting the Environment Variables :

| Variable Name | Description          | Default value |
| ------------- | -------------------- | ------------- |
| PORT          | Http port number     | `8000`        |
| DB            | Sqlite database path | `./stats`     |


## 🤔 How to use it
The Web server is deployed on Kubernetes cluster hosted in DigitalOcean.
The Url of the web server is `https://64.227.73.95` 

### `GET` - _`/`_

```shell
curl --request GET \
  --url 'https://64.227.73.95/?n1=3&n2=5&limit=15&str1=fizz&str2=buzz'
```

- Params :
  | Name  | Type   |
  | ----- | ------ |
  | N1    | Int    |
  | N2    | Int    |
  | Str1  | String |
  | Str2  | String |
  | Limit | Int    |

- Response :
  ```json
  ["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
  ```
### `GET` - _`/stats`_

```shell
curl --request GET \
  --url 'https://64.227.73.95/stats'
```
- Response :
  ```json
   {
    "params": {
        "n1": 3,
        "n2": 5,
        "limit": 15,
        "str1": "fizz",
        "str2": "buzz"
    },
    "hits": 1
  }
  ```
