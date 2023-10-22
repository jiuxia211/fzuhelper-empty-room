// Code generated by hertz generator.

package empty_room

import (
	"context"
	"fzuhelper-empty-room/biz/model/empty_room"
	"fzuhelper-empty-room/biz/pack"
	"fzuhelper-empty-room/biz/service"
	"fzuhelper-empty-room/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetEmptyRoom(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(empty_room.EmptyRoomRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	if req.Account == nil && req.Password == nil {
		// 当未传入账号密码时,使用默认账号密码
		defaultAccount := constants.DefaultAccount
		defaultPassword := constants.DefaultPassword
		req.Account = &defaultAccount
		req.Password = &defaultPassword
	}
	resp := new(empty_room.EmptyRoomResponse)
	empty_room, err := service.NewEmptyRoomService(ctx).GetRoom(req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.RoomName = empty_room
}