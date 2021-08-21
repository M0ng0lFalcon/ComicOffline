<template>
  <div class="home">
    <b-card no-body header="Comic Names" bg-variant="dark" text-variant="white">
      <b-list-group flush v-for="name in comicNames" :key="name">
        <b-list-group-item button variant="dark">{{ name }}</b-list-group-item>
      </b-list-group>
    </b-card>
  </div>
</template>

<script>
import { GetComicNames } from "@/Server/ComicSerevr.js";

export default {
  name: "Home",
  data() {
    return {
      comicNames: null,
    };
  },
  async mounted() {
    this.comicNames = await GetComicNames();
    console.log(this.comicNames);
  },
};
</script>
