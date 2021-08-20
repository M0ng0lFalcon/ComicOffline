<template>
  <div class="home">
    <b-list-group>
      <b-list-group-item button>咒术回战</b-list-group-item>
      <b-list-group-item button>咒术回战</b-list-group-item>
    </b-list-group>
    <b-button @click="GetComicNames">test</b-button>
  </div>
</template>

<script>
import GetComicNames from "../Server/ComicServer.js";

export default {
  name: "Home",
  date() {
    return {};
  },
  methods: {
    GetComicNames() {
      GetComicNames();
    },
  },
};
</script>
