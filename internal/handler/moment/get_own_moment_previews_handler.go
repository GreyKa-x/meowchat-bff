package moment

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/moment"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOwnMomentPreviewsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOwnMomentPreviewsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := moment.NewGetOwnMomentPreviewsLogic(r.Context(), svcCtx)
		resp, err := l.GetOwnMomentPreviews(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
