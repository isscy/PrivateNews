import http from "@/api/request";

export interface NewsDataEntity {
    id: number
    title: string;
    topic: string;
    tags: string;
    tagList: string[];
    initContent: string;
    content: string;
    international: string;
    internationalName: string,
    region: string,
    regionName: string,
    eventType: string,
    eventTypeName: string,
}

export interface ArticlePageSearchEntity {
    title: string;
    topic: string;
    content: string;
    international: string;
    region: string,
    eventType: string,
    currentPage: number
    pageSize: number
    searchValue: string
}

export interface PagedEntity4NewsData {
    currentPage: number
    pageSize: number
    totalCount: number
    pageCount: number
    rowDataList: NewsDataEntity[]
}




// 新增一篇文章
export function addArticle(data: NewsDataEntity) {
    return http<NewsDataEntity>({
        url: "/v1/publish",
        method: "post",
        data: data,
    });
}


// 分页获取文章列表
export function getArticlePage(data: ArticlePageSearchEntity) {
    return http<PagedEntity4NewsData>({
        url: "/v1/newsPage",
        method: "post",
        data: data,
    });
}