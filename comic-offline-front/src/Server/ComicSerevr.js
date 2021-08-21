import Vue from "vue";

export function GetComicNames() {
  return Vue.axios.get("/comic/comicNames").then((res) => {
    const data = res.data.data;
    const comicNames = data.comicNames;
    return comicNames;
  });
}
