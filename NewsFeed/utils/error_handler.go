package utils

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppErr struct {
	errMessage string
	statusCode codes.Code
}

func (e *AppErr) Error() string {
	return e.errMessage
}

func (e *AppErr) GetStatusCode() codes.Code {
	return e.statusCode
}

func (e *AppErr) SetErrMessage(err string) {
	e.errMessage = err
}

func (e *AppErr) SetStatusCode(code codes.Code) {
	e.statusCode = code
}

func ToGRPCStatus(err error) error {
	var appErr *AppErr
	if errors.As(err, &appErr) {
		return status.Error(appErr.statusCode, appErr.errMessage)
	}
	return status.Error(codes.Internal, err.Error())
}

var (
	ConnectingToDbError = &AppErr{
		errMessage: "error connecting to database",
		statusCode: codes.Internal}
	DatabaseQueryError = &AppErr{
		errMessage: "database query error",
		statusCode: codes.Internal}
	EmailAlreadyInUse = &AppErr{
		errMessage: "user with such email already exists",
		statusCode: codes.InvalidArgument}
	UsernameAlreadyInUse = &AppErr{
		errMessage: "user with such username already exists",
		statusCode: codes.InvalidArgument}
	ErrorGeneratingSaltForHashing = &AppErr{
		errMessage: "error hashing password",
		statusCode: codes.Internal}
	EqualIdsError = &AppErr{
		errMessage: "follower and followee ids are equal",
		statusCode: codes.InvalidArgument}
	AlreadyFollowingError = &AppErr{
		errMessage: "already follows",
		statusCode: codes.InvalidArgument}
	NotFollowingError = &AppErr{
		errMessage: "error trying to unfollow, not such relation exists",
		statusCode: codes.InvalidArgument}
	NoChangeAppliedError = &AppErr{
		errMessage: "no change is applied",
		statusCode: codes.InvalidArgument}
	UserNotFound = &AppErr{
		errMessage: "user with such username not found",
		statusCode: codes.NotFound}
	InvalidEncodedHashPasswordError = &AppErr{
		errMessage: "invalid encoded hash format",
		statusCode: codes.Internal}
	DecodingPasswordError = &AppErr{
		errMessage: "error decoding password",
		statusCode: codes.Internal}
	IncorrectPassword = &AppErr{
		errMessage: "incorrect password",
		statusCode: codes.PermissionDenied}
	MissingRequestFieldsError = &AppErr{
		errMessage: "missing request fields",
		statusCode: codes.InvalidArgument}
	InvalidRequestPayload = &AppErr{
		errMessage: "invalid request payload",
		statusCode: codes.InvalidArgument}
)
