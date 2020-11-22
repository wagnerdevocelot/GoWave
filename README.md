# Compreendendo Templates

![](https://cdn-images-1.medium.com/max/800/1*bC82PrnjL07hdznY5iscZA.png)

Um template nos permite criar um documento e em seguida, mesclar dados com ele.

Estamos aprendendo sobre Templates para que possamos criar um documento, uma página da web e, em seguida, mesclar dados personalizados a essa página. Templates Web permitem fornecer resultados personalizados aos usuários.

Pense no Twitter — você entra na sua página principal e vê os resultados personalizados para você. Essa página principal foi criada uma vez. É um template. No entanto, para cada usuário, esse template é preenchido com dados específicos desse usuário.

Outra exposição comum a Templates que a maioria de nós recebe todos os dias — Spam.

Uma empresa cria uma correspondência para enviar a todos e em seguida, mescla os dados com esse template para personalizar a correspondência para cada indivíduo. O resultado:

![](https://cdn-images-1.medium.com/max/800/1*GDYB6LMRnCaMug1ZT1_obA.png)

```
Caro Sr. Jones,

Você está cansado de contas altas de eletricidade?

Notamos que sua casa em …
```

```
Caro Sr. Smith,

Você está cansado de contas altas de eletricidade?

Notamos que sua casa em …
```

![](https://cdn-images-1.medium.com/max/800/1*qqceqAkGRkfe-I_pBjpNaw.png)

```
Caro {{Name}},

Você está cansado de contas altas de eletricidade?

Notamos que sua casa em …
```

Se você tiver alguma experiência com desenvolvimento web, você consegue olhar para uma página mesmo sem olhar o código fonte e identificar, o que é “template” e o que são “dados”.

Me faz lembrar também daqueles brinquedos de criança com geometrias e cores, onde o nosso texto seria a peça principal do brinquedo com os espaços específicos onde podemos inserir os objetos geométricos que seriam como os nossos dados.

![Brinquedo de encaixe de formas geométricas nas cores azul, verde, vermelho e amarelo](https://cdn-images-1.medium.com/max/800/1*ix-hWS3QOiguAinKdLBG_g.png)


![](https://cdn-images-1.medium.com/max/800/1*Ir_PP3OQJVPgVtv0qhgPPA.png)


Basicamente o que define o uso de templates em Go são duas ações, primeiro fazemos parse dos dados e depois executamos eles em alguma saída.

Parse:
```go
    tpl, err := template.ParseFiles("tpl.html")
    if err != nil {
        log.Fatal(err)
    }
```
Essa primeira etapa vai ser sempre muito parecida, você vai parsear esse conteúdo de algum arquivo ou vai receber ele como string de uma função, mas o destino é o mesmo, parsear.

Inclusive, o parâmetro de [_ParseFiles_](https://golang.org/src/text/template/helper.go?s=1224:1279#L27) é variado, então você pode parsear mais de um documento na mesma função ou você pode usar a variável tpl que foi criada pra fazer uma nova chamada assim _tpl.ParseFiles(novos argumentos)_.

Executando o conteúdo parseado em uma saída de terminal:
```go
    err = tpl.Execute(os.Stdout, nil)
    if err != nil {
        log.Fatal(err)
    }
```

Agora o segundo exemplo é genérico, pois como expliquei existem diversas opções de saída, você pode enviar o conteúdo parseado pra uma saída de terminal ou pra um arquivo por exemplo.

Você pode também usar _tpl.ExecuteTemplate(saída, nome, err)_ que permite não apenas definir uma saída como o exemplo acima mas também executar um template especifico usando o argumento do nome que seria uma string com o nome do documento em questão.

Então se você usou mais de um argumento em ParseFiles, pode ser que essa opção seja útil pra que você não tenha que executar todos os templates obrigatoriamente.

Agora já sabemos pra quê servem templates e como parsear e executar um ou mais templates. Vamos ver então um exemplo de como parsear mais de um template de forma otimizada.

![](https://cdn-images-1.medium.com/max/800/1*3l36sQSo4wWM1EHVdEbdkw.png)

[Arquivos](https://github.com/wagnerdevocelot/VaporWeb-GoWave/tree/master/perform)

Logo de inicio temos a importação dos pacotes de “_text/template_”, “_log_” para tratamento de erros e “_os_” para usar como saída padrão dos templates parseados.

```go
package main

import (
	"log"
	"os"
	"text/template"
)
```
Definimos tpl como uma variável com escopo global que aponta para o pacote de template.

```go
var tpl *template.Template
```

A função _init()_ é executada antes da _func main()_ O objetivo principal da função _init()_ é inicializar as variáveis ​​globais que não podem ser inicializadas no contexto global.

Aqui usamos [_template.Must_](https://golang.org/src/text/template/helper.go?s=576:619#L11) que nos auxilia com error handling, em seguida chamamos [_ParseGlob_](https://golang.org/src/text/template/helper.go?s=3809:3858#L93) com o path de uma pasta com os templates que desejamos parsear.

O ParseGlob é equivalente o ParseFiles, a diferença é que ele lê um path o e faz o parse dos documentos todos de uma vez.

Com essa combinação parseamos todos os arquivos dentro da pasta em questão de uma só vez e ai vamos distribuindo a saída uma a uma.

O “_nome_da_pasta/*_” significa que quero parsear todos os documentos dentro da pasta.

```go
func init() {
	tpl = template.Must(template.ParseGlob("nome_da_pasta/*"))
}
```

```go
func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "arquivo1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "arquivo2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "arquivo3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
```

A ordem de execução seria arquivo1, arquivo1, arquivo2, arquivo3, no caso o arquivo 1 retorna duas vezes pois ele é executado individualmente na função _tpl.ExecuteTemplate_ e na _tpl.Execute_ pois ele seria o primeiro na ordem dentro da pasta.

![](https://cdn-images-1.medium.com/max/800/1*soEy9d4wCaSova72Hn2s0w.png)

[Arquivos](https://github.com/wagnerdevocelot/VaporWeb-GoWave/tree/master/templateData)

Lembrando a questão do exemplo do e-mail, até o momento apenas fizemos o parse e a saída dos nossos templates, mas não os tornamos dinâmicos de fato, para isso precisaremos passar tipos de dados do Go dentro dos nossos templates.

E para isso dentro dos arquivos precisamos usar uma [_notação_](https://golang.org/pkg/text/template/#pkg-overview) com duas chaves, usando isso no momento do parse a função vai identificar essa notação e inserir nossos dados no template.

Em uma pasta “data” vou criar um arquivo chamado “dados_pessoais.txt” com esse conteúdo:

```text
Vamos testar aqui meu nome é: {{.Nome}}
Minha idade: {{.Idade}}
```

Quando a função ler esse arquivo e identificar as chaves duplas vamos ter dois dados que vem de uma struct. O ponto representa a struct em si e “Nome” sendo a propriedade da struct. Se fosse uma varável comum por exemplo, seria usado apenas o ponto, assim: ```{{.}}```

Se eu estivesse passando um slice eu poderia usar uma [_Action_](https://golang.org/pkg/text/template/#hdr-Actions) dentro do template:

```text
{{range .}}	{{.}} {{end}}
```

Assim você dá um range nos itens do slice o que se torna uma opção viável para listas em html.

A estrutura de um parse com dados vindos de código Go para dentro do arquivo fica então dessa forma:

```go
package main

import (
	"text/template"
	"log"
	"os"
)

type data struct {
	Nome  string
	Idade int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("data/*"))
}

func main() {

	perfil := data{
		"vapordev",
		29,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "dados_pessoais.txt", perfil)
	if err != nil {
		log.Fatalln(err)
	}
}
```

A estrutura é como a anterior, esse arquivo fica fora da pasta onde ficam os arquivos de texto com os templates a serem parseados, por isso as propriedades da struct estão com inicial em maiúscula, para que possam sem importadas para fora do pacote main.

As maiores diferenças aqui são que, instanciamos tipo data onde entrego a minha idade, e passo perfil como argumento na função.

Como seguimos o mesmo padrão, anterior é possível adicionar mais arquivos na pasta data e mais funções para a execução da saída.

A saída:
```text
Vamos testar aqui meu nome é: vapordev
Minha idade: 29
```

Obs. As vezes quando salvamos o arquivo, caso você mesmo não tenha inserido as importações e tenha deixado isso para o fmt fazer automaticamente é muito provável que ele importe o pacote errado, pois exite text/template e html/template os dois compartilham as mesmas funções com a diferença de que o pacote de template para html possui opções de segurança.

![](https://cdn-images-1.medium.com/max/800/1*slsz2_kofXUtkg99pqzPwQ.png)

[Arquivos](https://github.com/wagnerdevocelot/VaporWeb-GoWave/tree/master/twitter)

Agora que já temos certa ideia de como podemos parsear texto e usá-lo em uma saída vamos tentar usar esse conhecimento como representação de dados em uma página. Apesar de não estarmos utilizando html, ainda, já podemos pensar nessas representações, inclusive acho mais interessante pois as tags html a principio acabam poluindo o documento, com texto simples você tem menos ruído.

O exercício consiste em olhar uma página web do seu interesse, e tentar representar essa página ou pelo menos um pedaço dela em txt usando o pacote text/template para parsear coisas da página que a gente sabe que são dados varáveis e não apenas texto estático.

Como exemplo eu usei a página no twitter da [@tecnogueto](https://twitter.com/tecnogueto) com foco na parte de perfil.

![PrintScreen da página de perfil da Tecnogueto](https://cdn-images-1.medium.com/max/800/1*w8hzJibQ6e0n_xteMQx8NQ.png)


Aqui temos diversas informações que podemos parsear mas antes precisamos identificar quais seriam.

Pensando em tipos de dados, tudo que seria variável nessa página podemos já considerar como dados em Golang, como o numero de seguidores por exemplo.

A mídia da página, administradores podem modificar o **avatar** e a **capa**, assim também como a **bio**.

A palavra **Seguindo** e **Seguidores** por exemplo, é apenas texto isso não seria um dado Go, o numero de seguidores por outro lado sim.

O botão de **Seguir** e **Notificação** eu consigo ver como tipos **booleanos** enquanto que o Botão de **Direct** e **Opções** diria que se comportam como links então acredito que não seriam tipos de dados do Go exatamente.

Esse é um exercício interessante para se fazer quando se está aprendendo sobre desenvolvimento web pois vai te trazendo mais clareza sobre como são construídas as páginas, como são escolhidos os tipos de dados num geral.

Trouxe a página da [@tecnogueto](https://twitter.com/tecnogueto) como exemplo não por acaso, nesse momento ta rolando uma campanha de MatchFunding para o desenvolvimento de uma plataforma online de ensino em tecnologia pensada do gueto para o gueto.

Vou deixar um link para quem quiser saber mais sobre a iniciativa e puder contribuir, se a empresa em que você trabalha tem foco em ajudar a comunidade, principalmente focada em diversidade, essa é uma ótima oportunidade de ter uma ação concreta.

Colabore: [Aqui](https://benfeitoria.com/tecnogueto)

Voltando ao exercício, identifiquei os dados e então criei structs com os tipos que seriam usados.

```go
type midia struct {
	Avatar string
	Capa   string
}

type tweetConfig struct {
	Tweets      float64
	Notificação bool
	Follow      bool
	DataDaConta string
	Seguindo    float64
	Seguidores  float64
}

type perfil struct {
	Nome        string
	User        string
	Bio         string
	Localização string
	Site        string
	midia
	tweetConfig
}
```
Decidi que seria mais interessante compor os tipos, por isso temos _tweetConfig_ e media como tipos individuais mas que compõem o tipo perfil.

Em **midia** tudo é string, pois inclusive com html quando passamos um arquivo na tag o caminho ou link da mídia é uma string então aqui não muda.

**tweetConfig** Temos floats para numero de tweets, seguidores e etc. Notificações e Follow boleanos e DataDaConta eu imagino que seja usado um tipo como o timestamp que temos em schema de bancos de dados, aqui vou usar liberdade criativa (Gambiarra) pra que seja uma string.

O restante dos dados strings, em um db usaríamos VARCHAR com uma limitação de caracteres e etc, mas quando passamos os dados para Go podem ser strings.

Instanciar esses dados termina assim:

```go
	tecnoGueto := perfil{
		"Tecnogueto",
		"tecnogueto",
		`Nós da Tecnogueto, queremos socializar o cenário atual, ensinando profissões 
dentro da tecnologia e promovendo a diversidade de gênero, raça/etnia e cultural.`,
		"Rio de Janeiro, Brasil",
		"tecnogueto.com.br",
		midia{"https://pbs.twimg.com/profile_images/1162103759218139137/CYfSOQFx_400x400.jpg", "https://pbs.twimg.com/profile_banners/1002022507413757958/1553770868/1500x500"},
		tweetConfig{
			448,
			true,
			true,
			"Maio de 2018",
			100,
			3.2343,
		},
	}
```
Essa variável é o que vamos passar como argumento para a função de executar template.

Depois disso, parsear, executar, escolher uma saída é igual ao que já fizemos então vou deixar os detalhes a sua escolha.

O meu arquivo de perfil ficou assim:

```text
{{.Nome}}
{{.Tweets}} Tweets
{{.Capa}}

{{.Avatar}} Opções  Direct  {{.Notificação}} {{if .Follow}} Seguindo {{else}} Seguir {{end}}
{{.Nome}}
@{{.User}}

{{.Bio}}

{{.Localização}}    {{.Site}}  Ingressou em {{.DataDaConta}}

{{.Seguindo}} Seguindo  {{.Seguidores}} Seguidores
```

Como pode ver, além dos dados, usei uma condição como Action na propriedade de Follow, assim se instanciar como true aparece que estou seguindo, senão aparece Seguir.

Não imaginei como faria a mesma coisa com o botão de notificação pois ele não muda o texto mas sim o highlight, mas da pra pensar em algumas coisas.

A minha saída ficou assim:

```text
Tecnogueto
448 Tweets
https://pbs.twimg.com/profile_banners/1002022507413757958/1553770868/1500x500

https://pbs.twimg.com/profile_images/1162103759218139137/CYfSOQFx_400x400.jpg Opções  Direct  true  Seguindo 
Tecnogueto
@tecnogueto

Nós da Tecnogueto, queremos socializar o cenário atual, ensinando profissões 
dentro da tecnologia e promovendo a diversidade de gênero, raça/etnia e cultural.

Rio de Janeiro, Brasil    tecnogueto.com.br  Ingressou em Maio de 2018

100 Seguindo  3.2343 Seguidores
```

Obviamente, nada parecido com a página, a intenção nem é essa, mas sim entender melhor como is tipos Go se relacionam em um template. Quando falar mais sobre o pacote html/template as coisas vão ficar ainda mais interessantes.


![](https://cdn-images-1.medium.com/max/800/1*-XRRrXYvjbxv6RceHdTgJw.png)

Usar funções com dados em template pode ajudar com alguns tipos de tarefas e não seriam exatamente violações de “Separation of Concerns”, modificar a forma como os dados são apresentados é uma forma de se aplicar isso.

Uma função no template que vai modificar um dado lá no Data Base já é outra conversa.


![](https://cdn-images-1.medium.com/max/800/1*nsWgGlt59ppGf1pS9HWIMQ.png)

Podemos passar funções no template, vamos escolher aqui duas funções bem simples.

**_A primeira_**:

_Retorna os três primeiros caracteres de uma string passada em argumento, tranquilo._
```go
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}
```

**_A segunda_**:

_Uma função que já vem implementada no pacote string do Go._
```go
strings.ToUpper
```

![](https://cdn-images-1.medium.com/max/800/1*LPDK-a3qvE8fQVm-8iS5Iw.png)

Já conhecemos as funções, agora como elas são passadas no template?

Essa estrutura já é nossa velha conhecida, porém temos duas modificações, [**template.new**](https://golang.org/src/text/template/template.go?s=1022:1053#L27) inicia um template, o foco por enquanto é [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160) que vem antes de [**ParseFiles()**](https://golang.org/src/text/template/helper.go?s=2001:2070#L42) no method chaining.

```go
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
```
[**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160) tem um argumento do tipo [**FuncMap**](https://golang.org/src/text/template/funcs.go?s=1008:1043#L21) que é esse “fm”, vamos dar uma olhada nele mas precisamos entender algumas particularidades antes.

[**FuncMap**](https://golang.org/src/text/template/funcs.go?s=1008:1043#L21) recebe dois argumentos, uma string e uma interface vazia.

Interface vazia é uma interface sem nenhum método, um tipo Go tem pelo menos nenhum método, a interface vazia implementa todos os tipos Go, inclusive os tipos customizados que você quiser criar.

Então significa que [**FuncMap**](https://golang.org/src/text/template/funcs.go?s=1008:1043#L21) recebe como argumento uma string e qualquer coisa.

![](https://cdn-images-1.medium.com/max/800/1*2WpBuvei_6Dmqe9IMGKzaA.png)

Estamos passando uma string que vai funcionar como apelido das nossas funções e as funções em si que é o argumento da interface vazia que recebe qualquer coisa.

```go
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}
```
fm vai como argumento para [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160) e em seguida processado em [**ParseFiles()**](https://golang.org/src/text/template/helper.go?s=2001:2070#L42) que projeta os dados no template juntamente com as funções pré definidas.

A partir daqui é o básico, instanciar o objeto com os dados, passar como argumento para função de Execute com uma escolha de saída.

[**Arquivo**](https://github.com/wagnerdevocelot/GoWave/blob/master/template_funcs/main.go)

Atribuímos “apelidos” para as funções, passamos as funções como parametro na chaining de [**ParseFiles()**](https://golang.org/src/text/template/helper.go?s=2001:2070#L42).

Como decidimos quais dados irão receber essas funções? Usamos o apelido da função no template!

-   uc = Upper Case
-   ft = First three

```
{{uc .}} <!-- apelido + dado -->
```

![](https://cdn-images-1.medium.com/max/800/1*orG1YOaeb7TLpxdvbpqVAw.png)

Repare que estamos usando html, mas o pacote ainda é o text/template.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>functions</title>
</head>
<body>

{{range .}}
{{uc .Manufacturer}}
{{end}}


{{range .}}
{{ft .Manufacturer}}
{{end}}


</body>
</html>
```

Range passa por duas structs que estão dentro de um slice do qual usaremos as funções na propriedade Manufacturer, por isso o Range é aplicado somente ao “.” que é uma struct chamada car.

A nossa saída de template:

```text
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>functions</title>
</head>
<body>

FORD

TOYOTA


For

Toy

</body>
</html>
```

A primeira deixando os nomes das marcas em maiúsculas, e a segunda imprimindo somente os três primeiros caracteres.

![](https://cdn-images-1.medium.com/max/800/1*1cEGPl1Ksnlk5Plj3hbFmw.png) 

Já entendemos melhor o resultado, vamos analisar essa cadeia de métodos novamente pois é bem fácil se perder aqui.

```go
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
```

Talvez você se pergunte pra que serve o [**template.new**](https://golang.org/src/text/template/template.go?s=1022:1053#L27), se verificarmos a documentação, [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160)  tem um receptor com um ponteiro para Template então precisamos ter um template existente na cadeia de métodos antes de usar [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160).

[**template.new**](https://golang.org/src/text/template/template.go?s=1022:1053#L27) vai criar esse template vazio que vai ser usado como receptor em [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160)  que em seguida vai aplicar os “apelidos” das nossas funções no template vazio, para só então fazermos o ParseFiles.

![](https://cdn-images-1.medium.com/max/800/1*om_C7IEuDdk_eaU_O5EkHA.png)

Outro exemplo de aplicação de função que não viola o principio de “Separation of Concerns” é a utilização de funções para modificar os dados de Horário e Data, aplicamos com funções uma disposição diferente no template, mas não estamos de fato fuçando no dado real.

Exemplo da aplicação: [**clique aqui**](https://github.com/wagnerdevocelot/GoWave/tree/master/date_formatting)


Podemos usar também Pipelines para aplicar o mesmo dado a várias funções.

Levando em consideração ainda as funções do exemplo anterior de upercase e três caracteres:

```text
{{. | uc | ft }}
```

Se o dado utilizado fosse Fiat a nossa saída seria assim:

```text
Fiat    FIAT    Fia
```

Dependendo da ocasião ter funções que modificam a disposição dos dados de acordo com o usurário tendo como base uma condicional no template é uma opção comum.

![](https://cdn-images-1.medium.com/max/800/1*rYma0lQ4ee69ouEUqsbHgA.png)

Além das funções que podem ser passadas para o template com [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160) existem funções globais de template que podemos usar seus “apelidos” sem a necessidade de utilizar [**template.new**](https://golang.org/src/text/template/template.go?s=1022:1053#L27) e [**Funcs()**](https://golang.org/src/text/template/template.go?s=5031:5082#L160) pois essas funções já são globais no pacote de template.

Um exemplo simples é a função de index, podemos chamar cada item de um slice através do seu index ao invés de chamar todos de uma só vez com range.

Assim:
```html
{{index . 0}} <!-- chamando o item 0 do slice -->
{{index . 2}} <!-- chamando o item 2 do slice -->
{{index . 1}} <!-- chamando o item 1 do slice -->
```

Aqui na documentação você encontra uma lista com todas as funções globais para templates: [Clique aqui](https://golang.org/pkg/text/template/#hdr-Functions)

![](https://cdn-images-1.medium.com/max/800/1*Zo94UUz9-X2LhNNA3is92A.png)

É possivel modularizar páginas assim você tem vários arquivos especializados em uma parte especifica da sua página principal e então chamar esse pedaço onde for necessário. É um pouco parecido com as [Partials](https://guides.rubyonrails.org/layouts_and_rendering.html#using-partials) do Ruby on Rails.

Exemplo:

Eu quero uma página _index_, mas eu não quero colocar nela, o conteúdo do meu _footer_ e _minha lista de compras_ pois se apresentar no _index_ e estiver em um só arquivo vai se tornar confuso para dar manutenção quando essa pagina tiver uma quantidade muito grande de conteúdo.

Tem uma template golang com uma notation diferente, ```{{define “”}} {{end}}``` define e uma string que seria o nome desse template.


[footer.gohtml](https://github.com/wagnerdevocelot/GoWave/blob/master/templates_aninhados/templates/footer.gohtml)
```html
{{define "footer"}}
<footer>
  Algumas informações de copyright ou talvez alguma informação do autor?
</footer>
{{end}}
```


[list.gohtml](https://github.com/wagnerdevocelot/GoWave/blob/master/templates_aninhados/templates/list.gohtml)
```html
{{define "lista"}}
    <h2>Lista de compras</h2>

    <ul>
      <li>Café</li>
      <li>Pão Integral</li>
      <li>Manteiga</li>
    </ul>
{{end}}
```

Repare que o nome dado na string não é o mesmo do arquivo, pois quando chamar esses dois arquivos separados no index, vamos chamar pela string.

Assim:


[index.gohtml](https://github.com/wagnerdevocelot/GoWave/blob/master/templates_aninhados/templates/index.gohtml)
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Lista de Compras</title>
  </head>
  <body>
	{{template "lista"}}
  </body>
  {{template "footer"}}
</html>
```
Usamos a notation template seguida do nome do template.


A nossa saída:
```text
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Lista de Compras</title>
  </head>
  <body>
  
    <h2>Lista de compras</h2>

    <ul>
      <li>Café</li>
      <li>Pão Integral</li>
      <li>Manteiga</li>
    </ul>

  </body>
  
<footer>
  Algumas informações de copyright ou talvez alguma informação do autor?
</footer>

</html>
```

O arquivo main.go faz as chamadas normalmente, usando [**ParseGlob()**](https://golang.org/src/text/template/helper.go?s=3809:3858#L93) para parsear todos os templates de uma vez.


Caso queira consultar o arquivo: [clique aqui](https://github.com/wagnerdevocelot/GoWave/blob/master/templates_aninhados/main.go)


Caso seja necessário é possível passar dados dentro de templates aninhados, chamando o dado nas duas notations.

Definição de onde o dado vai aparecer no template
```html
{{define "footer"}}
	<p>Olá senhor {{.}} gostariamos muito de...</P>
{{end}}	
```
Chamada do template passando o dado:
```html
{{template "footer" .}}
```

![](https://cdn-images-1.medium.com/max/800/1*u6oLiOcFH-zFHsseIZJM4w.png)

Eu extrai o html da página de [**documentação**](http://www.golangbr.org/doc/) do Go, dei uma limpada nas coisas que não eram necessárias como o Go playground e etc.

Você pode baixar o arquivo: [aqui](https://github.com/wagnerdevocelot/GoWave/blob/master/exercicio/index.gohtml)

O objetivo:

Criar tipos de dados diferentes para
 - Titulos "H1,H2,H3" e etc
 - Links
 - Paragrafos
*Titulos*, *Links* e *Paragrafos* serão slices e cada um pertence a sua própria struct, no total três.

Cada titulo, link e Paragrafo deve:
- Ser chamado no template através do seu index no slice.

Onde houver divs:
- Precisa ser feito um template aninhado.

Nesse paragrafo aplique uma função a sua escolha: **Tutoriais por Daniel Mazza**
Sugestão:
- Inverter strings

Opcional:
-   Pesquisar sobre o pacote [**Time**](https://golang.org/pkg/time/) e passar no final do **body** uma função que retorna Dia/Mês/Ano

O pacote de templates termina aqui, porém existe ainda muito mais na documentação que eu recomendo a leitura.

No repositório desse artigo vou deixar uma pasta de respostas livre nos dois exercícios, você poderá clonar esse repo, colocar seu exercício nesse repositório com seu user na pasta, ex:

```text
├── exercicio
|	 └── respostas
|		 └── wagnerdevocelot
└── index.gohtml
```	

A trilha de web não termina aqui, falaremos sobre servers na próxima oportunidade.