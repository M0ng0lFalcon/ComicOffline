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

export function GetChapterPages(comicName, chapter) {
  const ComicNameParam = "ComicName=" + comicName;
  const ChapterParam = "Chapter=" + chapter;
  return Vue.axios
    .get("/comic/chapterPages?" + ComicNameParam + "&" + ChapterParam)
    .then((res) => {
      const data = res.data.data;
      let pageUrls = data.pageUrls;
      const pageCount = data.pageCount;

      // add base url to page urls
      for (let i = 0; i < pageUrls.length; i += 1) {
        pageUrls[i] = res.config.baseURL + pageUrls[i];
      }

      // sort the page urls
      pageUrls.sort(function (a, b) {
        const aJpg = a.split("/").pop();
        const bJpg = b.split("/").pop();
        const aNum = aJpg.split(".")[0];
        const bNum = bJpg.split(".")[0];

        return parseInt(aNum) - parseInt(bNum);
      });

      return {
        pageUrls: pageUrls,
        pageCount: pageCount,
      };
    });
}
