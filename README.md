# Image of the Day

Simple test Vue.js app to display images/memes.

## Data Generation
The app currently relies on a `src/assets/data.json` for the data that needs to be created first.

```json
{
    "Title": "Memes",
    "URLs":[
        {
         "URL": "https://picsum.photos/200/300",
         "URLHash": "e4f4a50ad707f41b6aad188699ede8dd",
         "PrimaryColorHex": "#fefefe",
         "SecondaryColorHex": "#b23909"
        }
    ]
}
```

`URLHash` is not currently used, but it could be used in a way to prevent showing the same image twice in a later iteration.

A small Go app can be used to generate this JSON file based on a list of image URLs (one URL per line):

```
go run generator/main.go urls.txt > src/assets/data.json
```

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
