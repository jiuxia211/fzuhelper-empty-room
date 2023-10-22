namespace go model

struct BaseResp {
    1: required i64 code,       // 请求返回的状态码
    2: required string msg,     // 返回的消息
}