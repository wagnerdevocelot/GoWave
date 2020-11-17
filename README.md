# Compreendendo Templates

![](https://cdn-images-1.medium.com/max/800/1*bC82PrnjL07hdznY5iscZA.png)

Um template nos permite criar um documento e em seguida, mesclar dados com ele.

Estamos aprendendo sobre Templates para que possamos criar um documento, uma página da web e, em seguida, mesclar dados personalizados a essa página. Templates Web permitem fornecer resultados personalizados aos usuários.

Pense no Twitter — você entra na sua página principal e vê os resultados personalizados para você. Essa página principal foi criada uma vez. É um template. No entanto, para cada usuário, esse template é preenchido com dados específicos desse usuário.

Outra exposição comum a Templates que a maioria de nós recebe todos os dias — Spam.

Uma empresa cria uma correspondência para enviar a todos e em seguida, mescla os dados com esse template para personalizar a correspondência para cada indivíduo. O resultado:

### Exemplo de template — mesclado com dados


Caro Sr. Jones,

Você está cansado de contas altas de eletricidade?

Notamos que sua casa em …

Caro Sr. Smith,

Você está cansado de contas altas de eletricidade?

Notamos que sua casa em …

### Exemplo de template

Caro {{Name}},

Você está cansado de contas altas de eletricidade?

Notamos que sua casa em …

Se você tiver alguma experiência com desenvolvimento web, você consegue olhar para uma página mesmo sem olhar o código fonte e identificar, o que é “template” e o que são “dados”.

Me faz lembrar também daqueles brinquedos de criança com geometrias e cores, onde o nosso texto seria a peça principal do brinquedo com os espaços específicos onde podemos inserir os objetos geométricos que seriam como os nossos dados.

![Brinquedo de encaixe de formas geométricas nas cores azul, verde, vermelho e amarelo](https://cdn-images-1.medium.com/max/800/1*ix-hWS3QOiguAinKdLBG_g.png)


## Parse e Execute


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

### Parse Glob

(Arquivos)[https://github.com/wagnerdevocelot/VaporWeb-GoWave/tree/master/perform]

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

## Actions e Notações

(Arquivos)[https://github.com/wagnerdevocelot/VaporWeb-GoWave/tree/master/templateData]

Lembrando a questão do exemplo do e-mail, até o momento apenas fizemos o parse e a saída dos nossos templates, mas não os tornamos dinâmicos de fato, para isso precisaremos passar tipos de dados do Go dentro dos nossos templates.

E para isso dentro dos arquivos precisamos usar uma [_notação_](https://golang.org/pkg/text/template/#pkg-overview) com duas chaves, usando isso no momento do parse a função vai identificar essa notação e inserir nossos dados no template.

Em uma pasta “data” vou criar um arquivo chamado “dados_pessoais.txt” com esse conteúdo:

```text
Vamos testar aqui meu nome é: {{.Nome}}
Minha idade: {{.Idade}}
```

Quando a função ler esse arquivo e identificar as chaves duplas vamos ter dois dados que vem de uma struct. O ponto representa a struct em si e “Nome” sendo a propriedade da struct. Se fosse uma varável comum por exemplo, seria usado apenas o ponto, assim: {{.}}

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

## Exercicio

(Arquivos)[https://github.com/wagnerdevocelot/VaporWeb-GoWave/tree/master/twitter]

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

Obviamente, nada parecido com a página, a intenção nem é essa, mas sim entender melhor como is tipos Go se relacionam em um template. Quando falar mais sobre o pacote html/template as coisas vão ficar ainda mais interessantes ^^