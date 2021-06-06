<template>
  <div class="bg-primary wrapper">
  <p><img class="main-img" :src="image"></p>
  <button class="bg-secondary" @click="selectedDataset++">Next</button>
</div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
    dataset: [],
    selectedDataset: 0
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
}

.bg-primary {
  color: v-bind('secondaryColor');
  background-color:  v-bind('primaryColor');
}

.bg-secondary {
  background-color:  v-bind('secondaryColor');
  color: v-bind('primaryColor');
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
