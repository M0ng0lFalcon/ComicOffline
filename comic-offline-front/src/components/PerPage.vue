<template>
  <div class="perPage">
    <b-img :src="pageUrls[pageNum - 1]" center></b-img>
  </div>
</template>

<script>
export default {
  name: "per-page",
  props: {
    pageNum: {
      type: Number,
      default: 1,
    },
    pageUrls: Array(),
  },
};
</script>
