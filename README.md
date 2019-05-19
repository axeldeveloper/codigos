# helper code 

```
Give examples
```
- python
- R
- go
- php
- c#
- javascript
- html


### Git

Crie um novo repositório no GitHub. Para evitar erros, não inicialize o novo repositório 
com README , licença ou .gitignore arquivos. Você pode adicionar esses arquivos depois 
que seu projeto foi empurrado para o GitHub


**Abra Git Bash**

Mude o diretório de trabalho atual para o projeto local.
Inicialize o diretório local como um repositório Git.

-----------------------------------------------------------

**git init**

Adicione os arquivos no seu novo repositório local. Isso os faz para o primeiro commit.

**git add**
Adiciona os arquivos no repositório local e os etapas para confirmação. 
Para inutilizar um arquivo, use 'git reset HEAD YOUR-FILE '.

**git commit -m "First commit"**
Confirma as alterações rastreadas e prepara-as para serem empurradas para um repositório 
remoto. Para remover este commit e modificar o arquivo, use 'git reset --soft HEAD ~ 1' e 
confirme e adicione o arquivo novamente.

No topo da página de configuração rápida do seu repositório GitHub, clique em  para copiar o URL do repositório remoto.
No prompt do Comando, adicione o URL do repositório remoto onde o seu repositório local será empurrado.

```
git remote add origin remote repository https://github.com/axeldeveloper/codigos.git
```



* [github Branches](https://git-scm.com/book/de/v2/Git-Branching-Remote-Branches) - Branches


**git remote -v**

SYNOPSIS

    - git-remote
    - git-remote add [-t <branch>] [-m <master>] [-f] [--mirror] <name> <url>
    - git-remote rm <name>
    - git-remote show <name>
    - git-remote prune <name>
    - git-remote update [group]


**git push origin master**
```
Empurra as alterações em seu repositório local até o repositório remoto que você especificou como origem
```


## Deployment
    Axel Alexander

## Built With

* [github](https://github.com/docs/) - The git repository

## Contributing



## Versioning



## Authors

* **Axel Alexander ** - *web site* - [PurpleBooth](http://axel-dev.herokuapp.com/)

 See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details