namespace go empty_room

include "model.thrift"

struct EmptyRoomRequest{
    1: required string time,
    2: required string start,
    3: required string end,
    4: required string building,
    5: optional string account,
    6: optional string password,
}

struct EmptyRoomResponse{
    1: required model.BaseResp base,
    2: required list<string> room_name,
}

service EmptyRoomService{
    EmptyRoomResponse GetEmptyRoom(1:EmptyRoomRequest req) (api.post="empty_room/query"),
}