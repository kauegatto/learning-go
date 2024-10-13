# Goroutines
V2 is V1, but with goroutines, when a blocking occurs, the green thread leaves the thread :)

In order to launch the green threads, you need to put the `go` keyword before a func, we can use an anonymous func for helping on that

```go
	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}
```
May seem to work at first, but note that since we're not waiting for the result to be completed, we're not really seeing changes correctly in the map.

Lets add a small timer:

```go
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}
```

Another existing problem is that we take each url from urls, and process `result[url] = wc(url)`. That seems alright, but note that each of our goroutines have a reference to the url variable - they don't have their own independent copy. So they're all writing the value that url has at the end of the iteration - the last url. Which is why the one result we have is the last url.

Vamos imaginar que temos um slice de URLs: urls := []string{"site1.com", "site2.com", "site3.com"}

O loop começa e cria uma goroutine para cada URL.
Cada goroutine tem uma referência à variável url, não uma cópia do valor.
O loop termina rapidamente, antes que as goroutines comecem a executar.
Quando o loop termina, a variável url contém "site3.com" (a última URL do slice).
Agora as goroutines começam a executar. Quando cada uma delas acessa a variável url, todas veem o mesmo valor: "site3.com".
Como resultado, todas as goroutines processam "site3.com", em vez de cada uma processar uma URL diferente.

O problema é que as goroutines não capturam o valor de url no momento em que são criadas. Em vez disso, elas mantêm uma referência à variável url, cujo valor muda durante o loop e termina com o último valor.

Vamos passar a url por parâmetro para a goroutine:

```go

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(curr string) {
			results[curr] = wc(curr)
		}(url)
	}

    time.Sleep(2 * time.Second)

	return results
}

```
Se você tiver sorte, vai ver o resultado correto
```
goos: linux
goarch: amd64
pkg: cv2
cpu: AMD Ryzen 9 5900X 12-Core Processor            
BenchmarkCheckWebsites-24              1        2000871421 ns/op
PASS
ok      cv2     4.005s
```
mas normalmente vai receber um erro:
`fatal error: concurrent map writes`

Isso acontece pois estamos escrevendo em uma estrutura de dados que não é `thread safe` de maneira concorrente

>This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.

> Go can help us to spot race conditions with its built in race detector. To enable this feature, run the tests with the race flag: `go test -race`.

# Channels
As you can see, communication on concurrency is a mess, we could use semaphores, synchronize threads or use thread-safe data structures, but instead, a more natural approach is using Channels, Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

Defining an output channel basically means that when the thing is processed, it's going through the output channel

```go
package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
```
Sincronização implícita:

A chave para entender este código é que as operações de canal em Go são bloqueantes por padrão.
Quando o segundo for tenta ler do resultChannel com `r := <-resultChannel`, ele bloqueará até que haja um resultado disponível no canal.

O segundo loop roda exatamente len(urls) vezes, garantindo que todos os resultados sejam coletados.
## Extra: Usando semáforos
Não é porquê channels existem na linguagem, que as alternativas que conhecemos antes não são válidas!
```go
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			status := wc(u)
			mu.Lock()
			results[u] = status
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	return results
}
```
Aqui estamos usando um mutex, semáforo que controla o acesso ao map! No caso, apenas uma thread por vez pode acessar a zona crítica: `wg.Add(1)`
Rodando: `go test -race`, temos:
```
PASS
ok      cv2     1.005s
```