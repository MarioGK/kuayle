package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kuayle/kuayle-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestTransferHTTPErrorReturnsMissingUsersAsStructuredDetails(t *testing.T) {
	e := echo.New()
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPost, "/api/workspaces/import", nil), recorder)

	err := transferHTTPError(ctx, &service.WorkspaceTransferError{
		Code: "WORKSPACE_IMPORT_MISSING_USERS", Message: "Missing users",
		MissingUsers: []string{"One@Example.test", "two@example.test"},
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, recorder.Code)
	require.JSONEq(t, `{
		"error": {
			"code": "WORKSPACE_IMPORT_MISSING_USERS",
			"message": "Missing users",
			"details": [
				{"field":"missing_user","message":"one@example.test"},
				{"field":"missing_user","message":"two@example.test"}
			]
		}
	}`, recorder.Body.String())
}

func TestTransferHTTPErrorReturnsSlugConflict(t *testing.T) {
	e := echo.New()
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPost, "/api/workspaces/import", nil), recorder)

	err := transferHTTPError(ctx, &service.WorkspaceTransferError{Code: "WORKSPACE_SLUG_TAKEN", Message: "Workspace slug is already taken"})

	require.NoError(t, err)
	require.Equal(t, http.StatusConflict, recorder.Code)
}
