# pos_go_weather
Sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin)

### Rota
Na rota /cep ao informar o parametro `{cep}` com um valor de CEP válido é realizado a busca em serviços como `viaCEP` e `brasilAPI`, com os endereços e regiões encontradas realizamos a busca em `api.weather.com` para usar dados de graus celsius e saber a temperatura atual em graus celsius, fahrenheit e kelvin.
```sh
/cep/{cep}
```

### cloud run
```sh
https://weather-5hvmq2vufa-uc.a.run.app/cep/83255000
```
