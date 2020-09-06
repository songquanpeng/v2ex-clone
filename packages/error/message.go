package error

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "invalid parameters",
	ERROR_EXIST_TAG:                "such tag name already exists",
	ERROR_NOT_EXIST_TAG:            "no such tag exists",
	ERROR_NOT_EXIST_POST:           "no such post exists",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "fail to check token",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token is outdated",
	ERROR_AUTH_TOKEN:               "fail to generate token",
	ERROR_AUTH:                     "something went wrong with token",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
