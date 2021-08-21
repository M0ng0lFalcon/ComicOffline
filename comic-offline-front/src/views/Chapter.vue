<template>
  <div class="chapter">
    <b-card
      no-body
      header="Comic Chapters"
      bg-variant="dark"
      text-variant="white"
    >
      <b-list-group flush v-for="name in chapters" :key="name">
        <b-list-group-item variant="dark">
          {{ name }}
        </b-list-group-item>
      </b-list-group>
    </b-card>
  </div>
</template>

<script>
import { GetComicChapters } from "../Server/ComicSerevr.js";

export default {
  name: "Chapter",
  data() {
    return {
      chapters: null,
    };
  },
  async mounted() {
    const comicName = this.$route.params.comicName;
    this.chapters = await GetComicChapters(comicName);
    console.log(this.chapters);
  },
};
</script>
