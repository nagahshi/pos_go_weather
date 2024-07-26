# pos_go_weather
Sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin)

### Setup local
Possua docker devidamente instalado em seu ambiente, necessita ainda de uma chave para API de consulta [WeatherAPI](http://weatherapi.com). Atualize o arquivo `docker-compose.yaml` com a chave na váriavel de ambiente `WEATHER_API_KEY`.
Colocando a aplicação no ar:
```sh
docker-compose up
```
*ps: certifique-se que a porta: 8080 esteja disponível*

### Uso
Na rota /cep ao informar o parametro `{cep}` com um valor de CEP válido é realizado a busca em serviços como [ViaCEP](https://viacep.com.br/) e [BrasilAPI](https://brasilapi.com.br), com os endereços e regiões encontradas realizamos a busca em [WeatherAPI](http://weatherapi.com) para usar dados de graus celsius e saber a temperatura atual em graus celsius, fahrenheit e kelvin.
```sh
/cep/{cep}
```

### cloud run
```sh
https://weather-5hvmq2vufa-uc.a.run.app/cep/83255000
```
