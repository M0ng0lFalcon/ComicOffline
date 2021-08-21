<template>
  <div class="ChapterPage">
    <b-pagination
      v-model="pageNum"
      :total-rows="rows"
      :per-page="perPage"
      aria-controls="page"
      align="fill"
      pills
    ></b-pagination>

    <b-card @click="nextPage">
      <per-page
        id="page"
        :pageNum="pageNum"
        :pageUrls="pagesInfo.pageUrls"
      ></per-page>
    </b-card>

    <b-pagination
      v-model="pageNum"
      :total-rows="rows"
      :per-page="perPage"
      aria-controls="page"
      align="fill"
      pills
    ></b-pagination>
  </div>
</template>

<script>
import { GetChapterPages } from "../Server/ComicSerevr.js";
import perPage from "../components/PerPage.vue";

export default {
  name: "ChapterPage",
  components: {
    perPage,
  },
  data() {
    return {
      pagesInfo: {
        pageCount: 0,
        pageUrls: [],
      },
      perPage: 1,
      pageNum: 1,
    };
  },
  methods: {
    nextPage() {
      this.pageNum += 1;
    },
  },
  async mounted() {
    const comicName = this.$route.params.comicName;
    const chapter = this.$route.params.chapter;
    this.pagesInfo = await GetChapterPages(comicName, chapter);
  },
  computed: {
    rows() {
      return this.pagesInfo.pageCount;
    },
  },
};
</script>
