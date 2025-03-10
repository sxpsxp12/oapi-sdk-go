// Package task code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/task/v1"
)

// PUT /open-apis/task/v1/tasks/:task_id/comments/:comment_id
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larktask.NewUpdateTaskCommentReqBuilder().
		TaskId("83912691-2e43-47fc-94a4-d512e03984fa").
		CommentId("6937231762296684564").
		UserIdType("user_id").
		Body(larktask.NewUpdateTaskCommentReqBodyBuilder().
			Content("飞流直下三千尺，疑是银河落九天").
			RichContent("飞流直下三千尺，疑是银河落九天<at id=7058204817822318612></at>").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Task.V1.TaskComment.Update(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
