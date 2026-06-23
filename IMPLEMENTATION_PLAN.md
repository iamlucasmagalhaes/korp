Parte 1: Criação do Serviço e Arquitetura do Ambiente
1.⁠ ⁠Serviço HTTP
•⁠  ⁠crie um servidor HTTP utilizando a linguagem Golang
•⁠  ⁠o serviço deve se chamar http-server-projeto-korp
•⁠  ⁠o serviço deve receber as requisições na porta 8080
•⁠  ⁠implemente um endpoint GET /projeto-korp
•⁠  ⁠esse endpoint deve retornar um JSON com a seguinte estrutura:
{
"nome": "Projeto Korp",
"horario": "<horário_atual>"
}
•⁠  ⁠O campo <horário_atual> deve conter o horário atual em UTC, resolvido
dinamicamente a cada requisição
•⁠  ⁠Crie um Dockerfile para a aplicação, que atenda aos seguintes requisitos:
•⁠  ⁠build
•⁠  ⁠execução da aplicação em container
2.⁠ ⁠Instalação e Configuração do Docker
•⁠  ⁠Em um ambiente Linux de sua escolha, instale e configure o Docker.
3.⁠ ⁠Configuração de Rede Docker
•⁠  ⁠Crie uma rede Docker no modo bridge para comunicação entre containers.
4.⁠ ⁠Docker Compose
Utilize Docker Compose para configurar dois containers:
● Container 1: http-server-projeto-korp
•⁠  ⁠Baseado na imagem construída anteriormente
•⁠  ⁠Conectado à rede criada
•⁠  ⁠Não deve expor portas diretamente ao host
● Container 2: nginx
•⁠  ⁠Imagem oficial do NGINX: https://hub.docker.com/_/nginx
•⁠  ⁠Conectado à mesma rede do serviço http-server-projeto-korp
•⁠  ⁠Porta 80 do host mapeada para a porta 80 do container
•⁠  ⁠Monte um volume no caminho /etc/nginx/conf.d/
5.⁠ ⁠Configuração do Proxy Reverso
No volume montado, adicione o arquivo http-server-projeto-korp.conf com a
configuração de proxy reverso
O NGINX deverá encaminhar requisições feitas à http://localhost:80 para o serviço na
porta 8080
6.⁠ ⁠Teste de Funcionamento
Teste o ambiente com o comando:
curl http://localhost:80/projeto-korp
A resposta esperada é o JSON gerado pelo http-server-projeto-korp
Parte 2: Monitoramento e Observabilidade
O objetivo dessa etapa é adicionarmos monitoramento ao serviço http-server-projeto-korp.
As seguintes métricas são obrigatórias:
•⁠  ⁠disponibilidade do serviço
•⁠  ⁠volume de requisições
A forma de expor a disponibilidade do serviço pode ser definida pelo candidato (ex: métrica,
endpoint dedicado, etc.). As métricas deverão ser expostas utilizando o padrão do
Prometheus.
Visualização das métricas
Altere o arquivo compose desenvolvido na Etapa 1 para que contenha os containers:
•⁠  ⁠Grafana
•⁠  ⁠Prometheus
Requisitos
•⁠  ⁠prometheus configurado para coletar as métricas expostas pelo serviço
•⁠  ⁠grafana configurado para visualizar essas métricas
•⁠  ⁠disponibilize um dashboard no Grafana que permita analisar o comportamento do
serviço
Parte 3: Automação com Ansible
O objetivo desse requisito é automatizar toda a configuração do ambiente descrita nas partes
1 e 2 usando Ansible.
Crie um playbook Ansible e ele deve contemplar, no mínimo:
•⁠  ⁠instalação do Docker em um ambiente Linux
•⁠  ⁠criação da rede Docker
•⁠  ⁠build da imagem do serviço http-server-projeto-korp
•⁠  ⁠criação e execução dos containers com docker compose
•⁠  ⁠configuração do NGINX com o proxy reverso
•⁠  ⁠configuração dos componentes de monitoramento
•⁠  ⁠validação do funcionamento do serviço desenvolvido através de uma requisição
HTTP, e posterior exibição da resposta no console
O ambiente deve ser totalmente provisionado com um único comando Ansible