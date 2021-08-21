import Vue from "vue";

export function GetComicNames() {
  return Vue.axios.get("/comic/comicNames").then((res) => {
    const data = res.data.data;
    const comicNames = data.comicNames;
    return comicNames;
  });
}

export function GetComicChapters(comicName) {
  return Vue.axios
    .get("/comic/comicChapters?ComicName=" + comicName)
    .then((res) => {
      const data = res.data.data;
      let chapters = data.Chapters;
      chapters.sort(function (a, b) {
        const aLi = a.split("-");
        const bLi = b.split("-");
        return parseInt(aLi[0]) - parseInt(bLi[0]);
      });
      return chapters;
    });
}
