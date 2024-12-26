package controllers

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/beego/beego/v2/server/web"
    "practice-project/controllers"
)

func setupTestContext(method, path string, body []byte) (*context.Context, *httptest.ResponseRecorder) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(method, path, bytes.NewBuffer(body))
    ctx := context.NewContext()
    ctx.Reset(w, r)
    return ctx, w
}

func TestShowCat(t *testing.T) {
    ctx, w := setupTestContext("GET", "/cat", nil)
    controller := &controllers.CatController{}
    controller.Init(ctx, "", "", nil)
    controller.ShowCat()
    assert.Equal(t, 200, w.Code)
}

func TestVoteUp(t *testing.T) {
    ctx, w := setupTestContext("POST", "/vote/up?image_id=test123", nil)
    controller := &controllers.CatController{}
    controller.Init(ctx, "", "", nil)
    controller.VoteUp()
    assert.Equal(t, 200, w.Code)
}

func TestCreateFavourite(t *testing.T) {
    request := controllers.FavouriteRequest{
        ImageID: "test123",
        SubID:   "user123",
    }
    body, _ := json.Marshal(request)
    
    ctx, w := setupTestContext("POST", "/favourites", body)
    controller := &controllers.CatController{}
    controller.Init(ctx, "", "", nil)
    controller.CreateFavourite()
    assert.Equal(t, 200, w.Code)
}