import http from "@/api/request";

export interface BulitInTagParams {
    tagKey: string;
    tagName: string;
    status: string
}

export interface BasePickerEntity {
    text: string;
    value: string;
}



// 获取内置Tag列表
export function getBulitInTagList(params: BulitInTagParams) {
    return http<BulitInTagParams>({
        url: "/v1/tag/getBuiltInTag",
        method: "get",
        params: params,
    });
}
