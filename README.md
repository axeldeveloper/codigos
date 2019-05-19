# helper code 
```
Give examples
```
python
go
php
c#
javascript
html



Reinitialized existing Git repository in


Crie um novo repositório no GitHub. Para evitar erros, não inicialize o novo repositório 
com README , licença ou gitignorearquivos. Você pode adicionar esses arquivos depois 
que seu projeto foi empurrado para o GitHub

# Abra Git Bash.

Mude o diretório de trabalho atual para o projeto local.
Inicialize o diretório local como um repositório Git.
-----------------------------------------------------------
git init
# Adicione os arquivos no seu novo repositório local. Isso os faz para o primeiro commit.

# git add .
Adiciona os arquivos no repositório local e os etapas para confirmação. 
Para inutilizar um arquivo, use 'git reset HEAD YOUR-FILE '.

# git commit -m "First commit"
Confirma as alterações rastreadas e prepara-as para serem empurradas para um repositório 
remoto. Para remover este commit e modificar o arquivo, use 'git reset --soft HEAD ~ 1' e 
confirme e adicione o arquivo novamente.


No topo da página de configuração rápida do seu repositório GitHub, clique em  para copiar o URL do repositório remoto.
No prompt do Comando, adicione o URL do repositório remoto onde o seu repositório local será empurrado.

git remote add origin remote repository https://github.com/axeldeveloper/codigos.git
# Define o novo remoto 
git remote -v
# Verifies the new remote URL
Pressione as alterações em seu repositório local para o GitHub.

git push origin master
# Empurra as alterações em seu repositório local até o repositório remoto que você especificou como origem