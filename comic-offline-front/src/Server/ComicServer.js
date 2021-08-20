import Vue from "vue";

export function GetComicNames() {
  return Vue.axios.get("/comic/comicNames").then((res) => {
    console.log(res);
  });
}
