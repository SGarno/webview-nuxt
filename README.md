# webview-nuxt

Testing zserge/webview and nuxtjs sample app

Assets from nuxt output are pre-compiled into assets.go, but to generate manually:

* Install nuxtjs
* Create sample app: ```npx create-nuxt-app sample```
* Generate distribution: ```npm run generate```
* Create assets.go file:  ```cd dist && go-bindata -o ..\assets.go ./...```


