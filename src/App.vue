<template>
<v-app>
  <div class="wrapper" :style="'background-color: '+primaryColor+' !important; color: '+secondaryColor+';'">
    <h1>{{title}}</h1>
  <p><v-img :src="image" class="main-img" transition="false"></v-img></p>
  <v-btn class="mt-2" size="x-large" @click="random()" :style="'background-color: '+secondaryColor+' !important; color: '+primaryColor+';'">{{buttonText}}</v-btn>
</div></v-app>
</template>

<script>

import dataset from './assets/data.json'
export default {
  name: 'App',
  data() {
    return {
    title: dataset.Title,
    dataset: dataset.URLs,
    selectedDataset: 0,
    left: []
    }

  },
  mounted() {
    this.random();
  },
  methods: {
    random() {
      if (this.left.length === 0) {
        for (let i = 0; i < this.dataset.length; i++) {
          this.left.push(i);
        }
      }
      let li = Math.floor(Math.random() * this.left.length);
      this.selectedDataset = this.left[li];
      this.left.splice(li,1);
    }
  },
  components: {
  },
  computed: {
    image() {
      return this.dataset[this.selectedDataset].URL
    },
    primaryColor() {
      return this.dataset[this.selectedDataset].PrimaryColorHex
    },
    secondaryColor() {
      return this.dataset[this.selectedDataset].SecondaryColorHex
    },
    buttonText() {
      if (this.left.length === 0) {
        return "Start Over"
      }
      return "Another!"
    }
  }
}
</script>

<style>
html,body,.wrapper {
        height:100%;
        margin:0;
        padding:0;
}

.main-img {
  height: 500px;
  max-height: 100%;
}

.wrapper {
  position: absolute;
  width: 100%;
  padding-top: 3em;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  
}
</style>
